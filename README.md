# Cloudflare R2 
Cloudflare is storage service that compatible with AWS S3. You can use AWS SDK to interact with this service. 

## Get Started
Before started, you need following package to be imported in your project:

```go 
import (
    "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)
```

After that, you can run below command to install its package and dependencies.
```sh
go mod tidy
```

## Configuration
Go to Cloudflare to grab your token. And then declare configuration using struct object below:

```go 
type R2Config struct {
    Bucket string 
    AccountID string 
    Key string 
    Secret string
}
```