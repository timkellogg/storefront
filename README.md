# Access Management System

Provides a way to authenticate users using oauth crendentials against applications.
The architecture is a API Gateway system


          Storefront
              |
    access management system (find out who the customer is, relay)
    |               |             
    orders        products
      |             |
    orders DB     products DB

storefront
  -> orders service -> orders DB
  -> product service -> products DB
  -> fulfillments service -> fulfillments DB


app structure

storefront -> UI
access management system
orders service
products service
messaging service
analytics service
fulfillments service
