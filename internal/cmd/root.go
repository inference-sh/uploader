package cmd

import (
	"github.com/spf13/cobra"
	"github.com/inference-sh/uploader/internal/upload"
)

func NewRootCmd() *cobra.Command {
	cfg := &upload.Config{}
	downloadCfg := &upload.DownloadConfig{}

	rootCmd := &cobra.Command{
		Use:   "uploader",
		Short: "Lightweight file uploader for S3-compatible storage",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.AddCommand(newS3Cmd(cfg))

	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download a file from a URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			return upload.Download(downloadCfg)
		},
	}

	downloadCmd.Flags().StringVar(&downloadCfg.URL, "url", "", "URL to download from")
	downloadCmd.Flags().StringVar(&downloadCfg.DestPath, "dest", "", "Destination file path")
	downloadCmd.MarkFlagRequired("url")
	downloadCmd.MarkFlagRequired("dest")

	rootCmd.AddCommand(downloadCmd)

	return rootCmd
}
