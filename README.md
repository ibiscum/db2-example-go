/usr/bin/bash '-c', 'docker run --rm -d -p 50000:50000/tcp icr.io/db2_community/db2:latest'

## Docker commands

Details at: https://www.ibm.com/docs/en/db2/11.5?topic=deployments-db2-community-edition-docker

Starting the DB2 container:
    docker run -h db2server --name db2server --restart=always --detach --privileged=true 
-p 50000:50000 --env-file .env_list -v icr.io/db2_community/db2:latest