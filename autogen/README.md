# AutoGen
AutoGen automatically generate test cases in FTW YAML format from Mod Security rule

## Status Tracker
```✅``` indicates that such function is implemented and ```⛔```️ means that we have no plan to support such function

### Operator

| Name                  	| Status 	| Used in CRS v3.2 	|
|-----------------------	|:------:	|:----------------:	|
| @beginsWith           	|    ✅   	|         ✅        |
| @contains             	|    ✅   	|         ✅        |
| @containsWord         	|        	|                  	|
| @endsWith             	|    ✅   	|         ✅        |
| @rsub                 	|        	|                  	|
| @rx                   	|    ✅   	|         ✅        |
| @pm                   	|    ✅   	|         ✅         |
| @pmFromFile           	|    ✅   	|         ✅         |
| @pmf                  	|        	|                  	|
| @streq                	|    ✅   	|         ✅         |
| @strmatch             	|    ✅   	|                  	|
| @within               	|    ✅   	|         ✅         |
| @eq                   	|    ✅   	|         ✅         |
| @ge                   	|    ✅   	|         ✅         |
| @gt                   	|    ✅   	|         ✅         |
| @le                   	|    ✅   	|                  	|
| @lt                   	|    ✅   	|         ✅         |
| @validateByteRange    	|    ✅   	|         ✅         |
| @validateDTD          	|        	|                  	|
| @validateHash         	|        	|                  	|
| @validateSchema       	|        	|                  	|
| @validateUrlEncoding  	|    ✅    	|         ✅         |
| @validateUtf8Encoding 	|    ✅   	|         ✅         |
| @detectSQLi           	|        	|         ✅         |
| @detectXSS            	|        	|         ✅         |
| @fuzzyHash            	|        	|                  	|
| @geoLookup            	|        	|                  	|
| @gsbLookup            	|    ⛔    	|                  	|
| @inspectFile          	|        	|                  	|
| @noMatch              	|    ⛔️   	|                  	|
| @rbl                  	|    ⛔️   	|        ✅          |
| @unconditionalMatch   	|    ⛔️   	|                  	|
| @verifyCC             	|        	|                  	|
| @verifyCPF            	|        	|                  	|
| @verifySSN            	|        	|                  	|
| @ipMatch              	|    ✅   	|        ✅          |
| @ipMatchFromFile      	|    ✅   	|        ✅          |
| @ipMatchF             	|        	|                  	|

### Variable

| Name                  	| Status 	| Used in CRS v3.2 	|
|-----------------------	|:------:	|:----------------:	|
| ARGS                  	|    ✅   	|         ✅        |
| ARGS_COMBINED_SIZE    	|    ✅    	|         ✅         |
| ARGS_NAMES            	|    ✅   	|         ✅         |
| ARGS_GET              	|    ✅   	|         ✅         |
| ARGS_GET_NAMES        	|    ✅    	|         ✅         |
| ARGS_POST             	|        	|                  	|
| ARGS_POST_NAMES       	|        	|                  	|
| FILES                 	|    ✅   	|                  	|
| FILES_COMBINED_SIZE   	|        	|         ✅         |
| FILES_NAMES           	|    ✅   	|         ✅         |
| FILES_SIZES           	|        	|                  	|
| FILES_TMPNAMES        	|        	|                  	|
| FILES_TMP_CONTENT     	|        	|                  	|
| FULL_REQUEST          	|        	|                  	|
| FULL_REQUEST_LENGTH   	|        	|                  	|
| MULTIPART_FILENAME    	|        	|                  	|
| MULTIPART_NAME        	|        	|                  	|
| PATH_INFO             	|        	|                  	|
| QUERY_STRING          	|    ✅    	|        ✅          |
| REMOTE_USER           	|        	|                  	|
| REQUEST_BASENAME      	|    ✅    	|        ✅          |
| REQUEST_BODY          	|    ✅   	|                  	|
| REQUEST_BODY_LENGTH   	|        	|                  	|
| REQUEST_COOKIES       	|    ✅   	|        ✅          |
| REQUEST_COOKIES_NAMES 	|    ✅   	|        ✅          |
| REQUEST_FILENAME      	|    ✅   	|        ✅          |
| REQUEST_HEADERS       	|    ✅   	|        ✅          |
| REQUEST_HEADERS_NAMES 	|    ✅   	|        ✅          |
| REQUEST_LINE          	|    ✅    	|        ✅          |
| REQUEST_METHOD        	|    ✅    	|        ✅          |
| REQUEST_PROTOCOL      	|    ✅    	|        ✅          |
| REQUEST_URI           	|    ✅    	|        ✅          |
| REQUEST_URI_RAW       	|    ✅    	|        ✅          |
| STREAM_INPUT_BODY     	|        	|                  	|

### Transformation

| Name                 	| Status 	| Used in CRS v3.2 	|
|----------------------	|:------:	|:----------------:	|
| t:base64Decode       	|    ✅   	|         ✅        |
| t:base64DecodeExt    	|        	|                  	|
| t:base64Encode       	|        	|                  	|
| t:cmdLine            	|        	|                  	|
| t:compressWhitespace 	|    ✅   	|         ✅         |
| t:cssDecode          	|        	|         ✅         |
| t:escapeSeqDecode    	|        	|                  	|
| t:hexDecode          	|    ✅   	|                  	|
| t:hexEncode          	|        	|         ✅         |
| t:htmlEntityDecode   	|        	|         ✅         |
| t:jsDecode           	|        	|         ✅         |
| t:length             	|    ✅   	|         ✅         |
| t:lowercase          	|    ✅   	|         ✅         |
| t:md5                	|        	|                  	|
| t:none               	|    ⛔️    	|         ✅         |
| t:normalisePath      	|    ✅   	|         ✅         |
| t:normalisePathWin   	|    ✅   	|         ✅         |
| t:normalizePath      	|    ✅   	|         ✅         |
| t:normalizePathWin   	|    ✅   	|         ✅         |
| t:parityEven7bit     	|        	|                  	|
| t:parityOdd7bit      	|        	|                  	|
| t:parityZero7bit     	|        	|                  	|
| t:removeComments     	|    ✅   	|         ✅         |
| t:removeCommentsChar 	|    ✅   	|                  	|
| t:removeNulls        	|    ✅   	|         ✅         |
| t:removeWhitespace   	|    ✅   	|                  	|
| t:replaceComments    	|    ✅   	|         ✅         |
| t:replaceNulls       	|    ✅   	|                  	|
| t:sha1               	|    ⛔️    	|                  	|
| t:sqlHexDecode       	|        	|                  	|
| t:trim               	|    ✅   	|         ✅         |
| t:trimLeft           	|    ✅   	|                  	|
| t:trimRight          	|    ✅   	|                  	|
| t:urlDecode          	|    ✅   	|         ✅         |
| t:urlDecodeUni       	|    ✅   	|         ✅         |
| t:urlEncode          	|        	|                  	|
| t:utf8toUnicode      	|        	|         ✅         |