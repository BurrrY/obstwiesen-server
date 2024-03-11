# obstwiesen-server



# TODO
- Adding Cultivars (Sorten)
- Adding Tags to everything

# Config

## Database-Connection
### SQLITE
- OW_DB_PROVIDER = sqlite
- OW_DB_CONNSTR = path to file (data.db)

### MySQL
- OW_DB_PROVIDER = mysql
- OW_DB_CONNSTR = root:my-secret-pw@tcp(127.0.0.1:3306)/meadows


## File-Storage
- OW_FILE_PROVIDER = "disk"
- OW_FILE_CONNSTR = "files" 

## GraphQL
- OW_GQL_PORT = 8080 //Port for the server to listen
- OW_GQL-PATH = /grapqh //Path where gql-server will be available

## Cross-Origin
- OW_XORIG = "http://localhost:13000"
- OW_XORIG_DBG = "true"

## Others
- OW_PUBLIC_URL = "http://localhost:8080"