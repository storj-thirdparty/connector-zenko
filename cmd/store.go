package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command to upload data to storj V3 network.",
	Long:  `Command to connect and transfer ALL tables from a desired Zenko instance to given Storj Bucket.`,
	Run:   zenkoStore,
}

func init() {

	// Setup the store command with its flags.
	rootCmd.AddCommand(storeCmd)
	var defaultZenkoFile string
	var defaultStorjFile string
	storeCmd.Flags().BoolP("accesskey", "a", false, "Connect to storj using access key(default connection method is by using API Key).")
	storeCmd.Flags().BoolP("share", "s", false, "For generating share access of the uploaded backup file.")
	storeCmd.Flags().StringVarP(&defaultZenkoFile, "zenko", "z", "./config/zenko_property.json", "full filepath contaning Zenko configuration.")
	storeCmd.Flags().StringVarP(&defaultStorjFile, "storj", "u", "./config/storj_config.json", "full filepath contaning storj V3 configuration.")
}

func zenkoStore(cmd *cobra.Command, args []string) {

	// Process arguments from the CLI.
	zenkoConfigfilePath, _ := cmd.Flags().GetString("zenko")
	fullFileNameStorj, _ := cmd.Flags().GetString("storj")
	useAccessKey, _ := cmd.Flags().GetBool("accesskey")
	useAccessShare, _ := cmd.Flags().GetBool("share")

	// Read Zenko instance's configurations from an external file and create an Zenko configuration object.
	configZenko := LoadZenkoProperty(zenkoConfigfilePath)

	// Read storj network configurations from and external file and create a storj configuration object.
	storjConfig := LoadStorjConfiguration(fullFileNameStorj)

	// Connect to storj network using the specified credentials.
	access, project := ConnectToStorj(fullFileNameStorj, storjConfig, useAccessKey)

	// Establish connection with Zenko and get io.Reader implementor.
	zenkoReader := ConnectToZenko(configZenko)

	// Fetch all file names and their corresponding Reader Object for upload
	uploadPathNames, zenkoObjectReaders := GetFileAndObjectNames(zenkoReader)

	fmt.Printf("\nInitiating back-up.\n")
	// Fetch all backup files from Zenko instance and simultaneously store them into desired Storj bucket.
	for i := 0; i < len(uploadPathNames); i++ {
		UploadData(project, storjConfig, uploadPathNames[i], zenkoObjectReaders[i])
	}
	fmt.Printf("\nBack-up complete.\n\n")

	// Create restricted shareable serialized access if share is provided as argument.
	if useAccessShare {
		ShareAccess(access, storjConfig)
	}
}
