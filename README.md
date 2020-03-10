# Shoping Cart Demo

Install steps coming soon!

# Carts Service
## Version: 1.0

**Contact information:**  
Mike Laeta  
https://laeta-shopping-cart-demo.firebaseapp.com  
mike@laetadevelopment.com  

### /v1/carts

#### POST
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [v1CreateRequest](#v1createrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1CreateResponse](#v1createresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/carts/all

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1ListResponse](#v1listresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/carts/{cart.id}

#### PUT
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cart.id | path |  | Yes | string |
| body | body |  | Yes | [v1UpdateRequest](#v1updaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1UpdateResponse](#v1updateresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

#### PATCH
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| cart.id | path |  | Yes | string |
| body | body |  | Yes | [v1UpdateRequest](#v1updaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1UpdateResponse](#v1updateresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/carts/{id}

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1ReadResponse](#v1readresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

#### DELETE
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1DeleteResponse](#v1deleteresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### Models


#### CartItem

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| title | string |  | No |
| qty | string |  | No |
| created | dateTime |  | No |
| updated | dateTime |  | No |

#### protobufAny

`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := ptypes.MarshalAny(foo)
     ...
     foo := &pb.Foo{}
     if err := ptypes.UnmarshalAny(any, foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| type_url | string | A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one "/" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading "." is not accepted).  In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:  * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.  Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics. | No |
| value | byte | Must be a valid serialized protocol buffer of the above specified type. | No |

#### runtimeError

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
| code | integer |  | No |
| message | string |  | No |
| details | [ [protobufAny](#protobufany) ] |  | No |

#### v1Cart

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| items | [ [CartItem](#cartitem) ] |  | No |
| created | dateTime |  | No |
| updated | dateTime |  | No |

#### v1CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| cart | [v1Cart](#v1cart) |  | No |

#### v1CreateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| id | string |  | No |

#### v1DeleteResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| deleted | string (int64) |  | No |

#### v1ListResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| data | [ [v1Cart](#v1cart) ] |  | No |

#### v1ReadResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| cart | [v1Cart](#v1cart) |  | No |

#### v1UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| cart | [v1Cart](#v1cart) |  | No |

#### v1UpdateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| updated | string (int64) |  | No |

# Items Service
## Version: 1.0

**Contact information:**  
Mike Laeta  
https://laeta-shopping-cart-demo.firebaseapp.com  
mike@laetadevelopment.com  

### /v1/items

#### POST
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [v1CreateRequest](#v1createrequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1CreateResponse](#v1createresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/items/all

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1ListResponse](#v1listresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/items/{id}

#### GET
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1ReadResponse](#v1readresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

#### DELETE
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| api | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1DeleteResponse](#v1deleteresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### /v1/items/{item.id}

#### PUT
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| item.id | path |  | Yes | string |
| body | body |  | Yes | [v1UpdateRequest](#v1updaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1UpdateResponse](#v1updateresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

#### PATCH
##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| item.id | path |  | Yes | string |
| body | body |  | Yes | [v1UpdateRequest](#v1updaterequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [v1UpdateResponse](#v1updateresponse) |
| 404 | Returned when the resource does not exist. | string (string) |
| default | An unexpected error response | [runtimeError](#runtimeerror) |

### Models


#### protobufAny

`Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := ptypes.MarshalAny(foo)
     ...
     foo := &pb.Foo{}
     if err := ptypes.UnmarshalAny(any, foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| type_url | string | A URL/resource name that uniquely identifies the type of the serialized protocol buffer message. This string must contain at least one "/" character. The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`). The name should be in a canonical form (e.g., leading "." is not accepted).  In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any. However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type URLs to message definitions as follows:  * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield a [google.protobuf.Type][]   value in binary format, or produce an error. * Applications are allowed to cache lookup results based on the   URL, or have them precompiled into a binary to avoid any   lookup. Therefore, binary compatibility needs to be preserved   on changes to types. (Use versioned type names to manage   breaking changes.)  Note: this functionality is not currently available in the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com.  Schemes other than `http`, `https` (or the empty scheme) might be used with implementation specific semantics. | No |
| value | byte | Must be a valid serialized protocol buffer of the above specified type. | No |

#### runtimeError

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
| code | integer |  | No |
| message | string |  | No |
| details | [ [protobufAny](#protobufany) ] |  | No |

#### v1CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| item | [v1Item](#v1item) |  | No |

#### v1CreateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| id | string |  | No |

#### v1DeleteResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| deleted | string (int64) |  | No |

#### v1Item

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |
| url | string |  | No |
| title | string |  | No |
| created | dateTime |  | No |
| updated | dateTime |  | No |

#### v1ListResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| data | [ [v1Item](#v1item) ] |  | No |

#### v1ReadResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| item | [v1Item](#v1item) |  | No |

#### v1UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| item | [v1Item](#v1item) |  | No |

#### v1UpdateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| api | string |  | No |
| updated | string (int64) |  | No |