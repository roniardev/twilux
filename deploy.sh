# remove all docker containers
docker rm -vf $(docker ps -aq)
docker rmi -f $(docker images -aq)
CM=~/deploy.sh

if test -f "$CM"; then
    rm -f ~/deploy.sh
fi

# create project dir and skip if it exist
mkdir -p twilux
cd twilux
# check if docker-compose.yml exist
FILE=~/twilux/docker-compose.yml
if test -f "$FILE"; then
    rm -f ~/twilux/docker-compose.yml
fi
# download my docker-compose.yml
wget https://gist.githubusercontent.com/enjinerd/09f552dfa935ce6198da8873c17c3271/raw/24b2947f48a85c7e8c6b5a91c84e3dcfb51631a9/docker-compose.yml
# build docker containers
docker-compose up -d
