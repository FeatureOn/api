docker build -t crdbinit crdbinit/.
docker run --network host crdbinit
