# jwt
Marshal and unmarshalling

1)marshalling means to convert a object(struct) into json body (byste string)
 and unmarshalling means to revert the byte string into struct
 
 Http Basic Authentication
● Basic authentication part of the specification of http
○ send username / password with every request
○ uses authorization header & keyword “basic”
■ put “username:password” together
■ converts them to base64
● puts generic binary data into printable form
● base64 is reversible
○ never use with http; only https
■ use basic authentication to login

Storing passwords
● never store passwords
● instead, store one-way encryption “hash” values of the password
● for added security
○ hash on the client
○ hash THAT again on the server
● hashing algorithms
○ bcrypt - current choice

 
