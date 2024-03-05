# obstwiesen-server



# TODO
- Adding Cultivars (Sorten)
- Adding Tags to everything

# Config

## Database-Connection
### SQLITE
- DB_PROVIDER = sqlite
- DB_CONNSTR = path to file (data.db)

### MySQL
- DB_PROVIDER = mysql
- DB_CONNSTR = root:my-secret-pw@tcp(127.0.0.1:3306)/meadows


## File-Storage
- PUBLIC_URL = "http://localhost:8080"
- FILE_PROVIDER = "disk"
- FILE_CONNSTR = "files" 