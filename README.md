# 📄 SMPP Protocol Error Code Definitions 

The SMPP (Short Message Peer-to-Peer) protocol is a widely-used industry standard for sending and receiving SMS traffic in the telecom industry. However, accessing the definitions for the various error codes used in the protocol can be inconvenient, as they are currently hosted in a PDF file on smpp.org.

To address this issue, i propose creating a new, free API service that provides updated definitions for the SMPP protocol error codes. This API is implemented in Go Lang, the definitions is stored in a managed MySQL database on Vultr cloud.

By providing easy access to these definitions, users will be able to easily update or correct their existing lists of error codes, or perform quick lookups on specific codes as needed. This will save time and effort for developers and professionals working with the SMPP protocol, and help ensure that they have the most accurate and up-to-date information at their disposal.

**Features**

    - Definitions for all SMPP protocol error codes
    - Easy-to-use API interface
    - Hosted on a reliable, high-performance cloud platform
    - Regularly updated to ensure accuracy

**Usag**e

To use the API, simply send a GET request to the API endpoint to get the update list of all SMPP error codes or with the error code you want to look up as a parameter. The API will return the definition for the specified error code in a JSON format.

For example:

```shell
curl https://URL:8083/errorcodes

[{"id":1,"command_status_name":"ESME_ROK","value":"0x00000000","description":"No error."},{"id":2,"command_status_name":"ESME_RINVMSGLEN","value":"0x00000001","description":"Message Length is invalid."}]

```

**Building the Binary**

To build the smpp_errors binary from the source code, you will need to have Go installed on your system.

Once you have Go set up, you can clone the repository and navigate to the root directory of the project:

```shell
$ git clone https://github.com/<username>/smpp-errors-api.git
$ cd smpp-errors-api
```

Next, build the binary by running the following command:

```shell

$ go build -o smpp_errors_api smpp_errors.go

```

This will create a binary file called smpp_errors_api in the root directory of the project. You can then run the binary by typing:

```shell

./smpp_errors_api

```

This will start the API service, which you can then access by sending requests to the specified endpoint.


**Contributing**

We welcome contributions to the API service, whether in the form of code, documentation, or suggestions for improvement. If you are interested in contributing, please feel free to open a pull request or issue on the GitHub repository.
License

This API service is provided under an MIT license. Feel free to use it as you see fit, but please be sure to give appropriate credit and follow the terms of the license.
