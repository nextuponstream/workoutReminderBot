FROM mongo:3.6.22
# FIXME doesn't copy when inspecting with
# docker exec -it wrb_mdb /bin/bash
# FIXME system prune outlined a new error: runBot shows 
# "bad address wrb_mdb"
COPY ./initdb/ /docker-entrypoint-initdb.d/