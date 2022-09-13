package http

import (
	"context"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-challenege/app_context"
	v1 "go-challenege/cmd/http/v1"
	"go-challenege/pkg/database"
	"go-challenege/pkg/httpserver"
	"go-challenege/pkg/i18n"
	"go-challenege/pkg/logger"
)

const (
	DbUri = "db-uri"

	ServMode             = "serv-mode"
	ServPort             = "serv-port"
	ServSupportLanguages = "server-support-languages"

	AccessTokenExpiry = "access-token-expiry"
	AccessTokenSecret = "access-token-secret"

	RefreshTokenExpiry = "refresh-token-expiry"
	RefreshTokenSecret = "refresh-token-secret"
)

func registerFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(DbUri, "", "")
	cmd.PersistentFlags().String(ServPort, "", "")
	cmd.PersistentFlags().String(ServMode, "", "")

	cmd.PersistentFlags().String(AccessTokenExpiry, "", "")
	cmd.PersistentFlags().String(AccessTokenSecret, "", "")

	cmd.PersistentFlags().String(RefreshTokenExpiry, "", "")
	cmd.PersistentFlags().String(RefreshTokenSecret, "", "")

	cmd.PersistentFlags().StringSlice(ServSupportLanguages, []string{"en", "vi"},
		"server language support when response")

	_ = viper.BindPFlag(DbUri, cmd.PersistentFlags().Lookup(DbUri))
	_ = viper.BindPFlag(ServPort, cmd.PersistentFlags().Lookup(ServPort))
	_ = viper.BindPFlag(ServMode, cmd.PersistentFlags().Lookup(ServMode))

	_ = viper.BindPFlag(AccessTokenExpiry, cmd.PersistentFlags().Lookup(AccessTokenExpiry))
	_ = viper.BindPFlag(AccessTokenSecret, cmd.PersistentFlags().Lookup(AccessTokenSecret))

	_ = viper.BindPFlag(RefreshTokenExpiry, cmd.PersistentFlags().Lookup(RefreshTokenExpiry))
	_ = viper.BindPFlag(RefreshTokenSecret, cmd.PersistentFlags().Lookup(RefreshTokenSecret))

	_ = viper.BindPFlag(ServSupportLanguages, cmd.PersistentFlags().Lookup(ServSupportLanguages))
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serv",
		Short: "Start HTTP Service",
		Long:  "Start HTTP Service",
		RunE: func(cmd *cobra.Command, args []string) error {
			var l = logger.Init(
				logger.WithLogDir("logs/"),
				logger.WithDebug(true),
				logger.WithConsole(true),
			)
			defer l.Sync()
			ctx := context.Background()
			dbCnf := database.MongoConfig{
				Uri: viper.GetString(DbUri),
			}

			mgoDB := database.NewAppDB(&dbCnf)
			mgoDB.Start(ctx)

			httpCnf := httpserver.NewMyHttpServerConfig(viper.GetString(ServMode), viper.GetString(ServPort))
			appI18n, _ := i18n.NewI18n(i18n.NewI18nConfig(viper.GetStringSlice(ServSupportLanguages)))
			router := httpserver.New(httpCnf, appI18n)

			sc := app_context.NewAppCtx(mgoDB.GetDB())

			router.AddHandler(v1.SetupRoute(sc))
			router.Start()

			gracehttp.Serve(router.Server)
			return nil
		},
	}
	return cmd
}

func RegisterCommandRecursive(parent *cobra.Command) {
	registerFlags(parent)
	cmd := NewServerCommand()
	parent.AddCommand(cmd)
}
