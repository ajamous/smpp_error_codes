# SMPP Protocol Error Code Definitions

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

**Contributing**

We welcome contributions to the API service, whether in the form of code, documentation, or suggestions for improvement. If you are interested in contributing, please feel free to open a pull request or issue on the GitHub repository.
License

This API service is provided under an MIT license. Feel free to use it as you see fit, but please be sure to give appropriate credit and follow the terms of the license.
