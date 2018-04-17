# TantanDemo
a micro go program to design TantanDemo

## summary
*  use beego web application to finish the api service. 
*  mvc frame. folder router for route, controller for calling service, model for creating tabel or operate data 

## design

*  url design.  
    + according to requirements design resful  api style

* db design
    + use unique union key constraint the id_state to find somebody state(like,dislike,match) quickly  and id_id to find 1 v 1 relation quickly 
    + maybe next step we can use transaction to deal state "match" problem 
    
