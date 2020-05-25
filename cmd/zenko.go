package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go"
)

// ConfigZenko defines the variables and types.
type ConfigZenko struct {
	EndPoint        string `json:"zenkoEndpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
}

// ZenkoReader implements an io.Reader interface
type ZenkoReader struct {
	Client *minio.Client
}

// LoadZenkoProperty reads and parses the JSON file
// that contains a Zenko instance's property
// and returns all the properties as an object.
func LoadZenkoProperty(fullFileName string) ConfigZenko { // fullFileName for fetching database credentials from  given JSON filename.

	var configZenko ConfigZenko

	// Open and read the file.
	fileHandle, err := os.Open(filepath.Clean(fullFileName))
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(fileHandle)
	if err = jsonParser.Decode(&configZenko); err != nil {
		log.Fatal(err)
	}

	if err = fileHandle.Close(); err != nil {
		log.Fatal(err)
	}

	// Display the read Zenko configuration properties.
	fmt.Println("Read Zenko configuration from the ", fullFileName, " file")
	fmt.Println("ZenkoEndPoint\t", configZenko.EndPoint)
	fmt.Println("AccessKeyID\t", configZenko.AccessKeyID)
	fmt.Println("SecretAccessKey \t", configZenko.SecretAccessKey)

	return configZenko
}

// ConnectToZenko will connect to a Zenko instance,
// based on the read property from an external file.
// It returns a reference to an io.Reader with Zenko instance information
func ConnectToZenko(configZenko ConfigZenko) *ZenkoReader {

	fmt.Println("\nConnecting to Zenko.")

	// Initialize minio client object.
	minioClient, err := minio.New(configZenko.EndPoint, configZenko.AccessKeyID, configZenko.SecretAccessKey, true)
	if err != nil {
		log.Fatalf("Failed to establish connection with Zenko: %s\n", err)
	}

	// Return Zenko connection client.
	return &ZenkoReader{Client: minioClient}
}

// GetFileAndObjectNames will return all file names
// and their corresponding Reader Object for upload
func GetFileAndObjectNames(zenkoReader *ZenkoReader) ([]string, []*minio.Object) {

	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true

	// List of all buckets from Zenko orbit.
	buckets, err := zenkoReader.Client.ListBuckets()
	if err != nil {
		log.Fatal("List Bucket Error:", err)
	}

	uploadFilePath := make([]string, 0)
	objectReaderArray := make([]*minio.Object, 0)

	t := time.Now()
	timeNow := t.Format("2006-01-02_15_04_05")

	for _, zenkoBucket := range buckets {
		// ListObjects lists all objects from the specified bucket.
		objectCh := zenkoReader.Client.ListObjects(zenkoBucket.Name, "", isRecursive, doneCh)
		for object := range objectCh {
			if object.Err != nil {
				log.Fatal("Object Information Error", object.Err)
			}

			if object.Key[len(object.Key)-1] != '/' {

				fmt.Println("\nReading from the file :", object.Key)
				zenkoFilePath := zenkoBucket.Name + "_" + timeNow + "/" + object.Key

				// GetObject function returns seekable, readable object.
				objectReader, err := zenkoReader.Client.GetObject(zenkoBucket.Name, object.Key, minio.GetObjectOptions{})
				if err != nil {
					log.Fatal(err)
				}

				uploadFilePath = append(uploadFilePath, zenkoFilePath)
				objectReaderArray = append(objectReaderArray, objectReader)

			}
		}
	}

	return uploadFilePath, objectReaderArray
}
