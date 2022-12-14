package internal

import (
    "context"
    "github.com/chatpuppy/puppychat/app/cache"
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/logger"
    "github.com/chatpuppy/puppychat/pkg/tea"
    log "github.com/sirupsen/logrus"
)

func ApiBootstrap() {
    // load configure from file
    g.InitializeConfig()

    // logger
    logger.LoadWithConfig(logger.Config{
        Color:  true,
        Level:  "info",
        Age:    8192,
        Caller: g.Config.App.Mode == "development",
    })

    // auto sync time
    go g.AutoSyncTime()

    // connect database
    ent.OpenDatabase(g.Config.Database.Dsn, g.Config.Database.Debug)

    // create cache client
    cfg := g.Config.Redis
    cache.CreateClient(cfg.Address, cfg.Password, cfg.DB)

    // cache all groups
    cacheMembers()

    log.Info("server bootstrap done")
}

// cache all group's member on start
func cacheMembers() {
    if cache.Group.Initialized {
        return
    }
    gms, _ := ent.Database.GroupMember.Query().All(context.Background())
    groups := make(map[string]*[]string)
    for _, gm := range gms {
        m, ok := groups[gm.GroupID]
        if !ok {
            m = tea.Pointer([]string{})
            groups[gm.GroupID] = m
        }
        *m = append(*m, gm.MemberID)
    }
    for g, items := range groups {
        cache.Group.Initialize(g, *items)
    }
    cache.Group.Initialized = true
}
