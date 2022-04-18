Server=$1

function check_parameter ()
{
    Param=$2
    ParamLen=${#Param}

    if [ -z $Param ] || [ $ParamLen -lt 1 ];
    then
        echo "Parameter $1 not passed"
        exit 1
    fi
}

check_parameter 'Server' $Server

mkdir /apps
mkdir /apps/connection

cd /apps/connection

if [ $Server -gt 0 ];
then
  echo Downloading files
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/server/server ./
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/server/config.yml ./
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/server/connectionserver.service ./

  echo Setting up service
  mv ./connectionserver.service /etc/systemd/system/
  systemctl daemon-reload
  service connectionserver start
else
  echo Downloading files
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/client/client ./
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/client/config.yml ./
  gcloud alpha storage cp gs://genuine-citron-347208-consul/api/client/connectionclient.service ./

  echo Setting up service
  mv ./connectionclient.service /etc/systemd/system/
  systemctl daemon-reload
  service connectionclient start
fi

chmod 750 /apps/connection/*

firewall-cmd --zone=public --add-port=9000/tcp --permanent
firewall-cmd --complete-reload

yum install nano -y