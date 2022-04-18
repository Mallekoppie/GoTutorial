set GOARCH=amd64
set GOOS=linux

echo Building Server
cd ./Server
go build -o "server"

echo Building Client
cd ../Client
go build -o "client"

echo Uploading to Cloud Storage
cd ..

echo Uploading Server
cmd /c gcloud alpha storage cp ./Server/server gs://genuine-citron-347208-consul/api/server/server
cmd /c gcloud alpha storage cp ./Server/config.yml gs://genuine-citron-347208-consul/api/server/config.yml
cmd /c gcloud alpha storage cp ./Server/connectionserver.service gs://genuine-citron-347208-consul/api/server/connectionserver.service

echo Uploading Client
cmd /c gcloud alpha storage cp ./Client/client gs://genuine-citron-347208-consul/api/client/client
cmd /c gcloud alpha storage cp ./Client/config.yml gs://genuine-citron-347208-consul/api/client/config.yml
cmd /c gcloud alpha storage cp ./Client/connectionclient.service gs://genuine-citron-347208-consul/api/client/connectionclient.service