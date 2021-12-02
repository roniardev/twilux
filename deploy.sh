# remove all docker containers
docker rm -f $(docker ps -a -q)
# create project dir and skip if it exist
mkdir -p twilux
cd twilux
# check if docker-compose.yml exist
FILE=~/twilux/docker-compose.yml
if test -f "$FILE"; then
    rm -f ~/twilux/docker-compose.yml
fi
# download my docker-compose.yml
wget https://gist.githubusercontent.com/enjinerd/09f552dfa935ce6198da8873c17c3271/raw/bdb1d511487ea100ac66dae19f542e8eb4164cce/docker-compose.yml
# build docker containers
docker-compose up -d