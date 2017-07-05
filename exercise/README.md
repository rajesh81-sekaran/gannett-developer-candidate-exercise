Overview
The code is written in golang. The code is designed as per the requirement.

Compilation:
The code can be compiled from this root directory where the server.go file is present

Command to use:
go build server.go

The above command will generate the server executable "server" in the current directory where the server.go file is located.

The port number on which the server listens, the content url and the profile url can be configured as environment variables. If these environment variables are not set these values are configured as constants in common/data.go file and these predefined values will be used. The default port on which the server will start is 9090.

SERVER_PORT
This environment variable defines the port number on which the server will listen

PROFILE_URL
Profile url that will be used to fetch the profile

CONTENT_URL
Content url which will be used to fetch the content

Example usage:
export SERVER_PORT=9091
export PROFILE_URL="https://peaceful-springs-7920.herokuapp.com/profile/"
export CONTENT_URL="https://peaceful-springs-7920.herokuapp.com/content/"

Once the above environment variables are set in a shell, the server process should be started(./server) from the same shell where these variables are exported. Once the server starts it will print these environment variables
E.g:
# pwd
/vagrant/GoPath/src/exercise
# ls -lrt
total 5
-rwxrwxrwx. 1 vagrant vagrant 1003 Jul  5 00:09 server.go
drwxrwxrwx. 1 vagrant vagrant 4096 Jul  5 00:11 common
-rwxrwxrwx. 1 vagrant vagrant    0 Jul  5 00:13 README.md
# go build server.go
# ls -lrt
total 8989
-rwxrwxrwx. 1 vagrant vagrant    1003 Jul  5 00:09 server.go
drwxrwxrwx. 1 vagrant vagrant    4096 Jul  5 00:11 common
-rwxrwxrwx. 1 vagrant vagrant       0 Jul  5 00:13 README.md
-rwxrwxrwx. 1 vagrant vagrant 9199224 Jul  5 00:24 server
# export SERVER_PORT=9091
# export PROFILE_URL="https://peaceful-springs-7920.herokuapp.com/profile/"
# export CONTENT_URL="https://peaceful-springs-7920.herokuapp.com/content/"
# ./server
Starting Server on Port: 9091
Profile URL: https://peaceful-springs-7920.herokuapp.com/profile/
Content URL: https://peaceful-springs-7920.herokuapp.com/content/

Once the server is started the client can connect to the ip address on which the server is started

E.g
http://192.168.33.10:9091

After the above step the client will be presented with a webpage to input the user name. Once the user name is provided by the client and the page is submitted, the server will fetch the data from the profile url and content url and present it to the client. The user name provided by the client will be remembered by the server and a prfile ID will be associated with this user name and this mapping will be retained by the server for 365 days. This mapping is done in the internal memoryof the server process to make the design simple.

Code organization:
server.go:
This is the code that has the main() function. The http handler function is initialized in this code. The servers listen function is also initialized in this code. This is the code that when compiled gnerates the server executable.

common:
The folder wich has the common function and the constant values

common/handler.go:
This code has the inmplementation of the http handler function and the business logic.
Functionality:
(1) Registers the http listener
(2) Serves the client with a webpage to input the user name
(3) Once the user name is provide by the client, the server gets the profile id using the profile url and maps this profile url to the user name and retains this mapping in the memmory for 365 days.
(4) After getting the profile name from the profile url, the content corresponding to the profile name is fetched from the content url.
(5) With this content an webpage is created with hyperlinks and fonts with appropriate text color.
(6) Once the web page is formed, it is then served to the client.

common/data.go:
This file has the constant values which will be used in the code. By modifying the values in this file the behaviour of the server can be customized.
Constants:
(1) DefaultPort : Defines the port on which the server starts. The port can also be configured using the environment variable "SERVER_PORT". The environment variable takes the precedence.
(2) DefaultProfileURL : Defines the url from which the profile will be fetched. This url can also be configured using the environment variable "PROFILE_URL". The environment variable takes the precedence.
(3) DefaultContentURL : Defines the url from which the contents will be fetched. This url can also be configured using the environment variable "CONTENT_URL". The environment variable takes the precedence.
Apart from the above constant, this file has other common structures and fucntions used in the code.

common/user.html:
This is a simple html page that has an input text field used by the client to input the user name, a submit button and a script to validate the user.
The following validations are done on the user field:
(1) User field is checked for the empty values
(2) Checking is done whether the first character of the user name starts only with a letter.

Following are the requirements for a valid user:
(1) The user name should not be empty.
(2) The user name should begin with a character/letter.

common/data.html:
This html file has the contents of the web page. It has the following entries
(1) Page headline "My Delicious Articles"
(2) Title with the hyperlink
(3) Summary directly below the Title. Summary is also populated with the hyperlink

The font color of the text depends on the "theme" value specific to the profile
If the theme is "rare" font color will be crimson (#DC143C).
If the theme is "well" font color will be saddle brown (#8B4513).

Packages used:
(1) Dynamic webpages are generated using the template package in golang(html/template).
(2) Handling of web requests and replies are handled by the built in http package(net/http).
(3) Synchronization is acheived by using the sync package(sync).
(4) Profile ID persistence is handled by time package(time).
(5) Parsing of json fields from the profile url and content url is handled by json package(encoding/json).
(6) Some of the string functions are handeld by strings package(strings).


Testing:
curl can be used to test the performance of the server

I have used the following curl command to test the server endpoint
curl -X POST -F 'user=SomeUser' -s http://192.168.33.10:9091?[1-1000] -o  /dev/null

Other tools like "ab" from Apache can be used to invoke concurrent requests and load test the server end point.