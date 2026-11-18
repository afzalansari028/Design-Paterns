# System Design

1.Functional vs non functional:
 
  Functionl : What feaatures system shoud provide.

  For an e-commerce system, examples are:
  
      User can register/login
      User can browse products 
      User can search products 
      User can add products to cart
      User can place an order
      User can make payment
      User can track an order
      Admin can add/update products

  Non-Functional : How the system should behave.

      These are things like:
      Scalability
      Availability
      Reliability
      Performance
      Low latency
      Security
      Consistency
      Durability

      Product search should respond within 200 ms.
      System should support 1 million users.
      System should remain available even if one server crashes.
      
Load Balancer: 

    - redirect the traffic
    - able to checkk the health of pods.
    
CDN - Content Delivery Network:

    A CDN is a world-wide network of servers distributed across different geographic locations that delivers content to users
    from a server which closer to them, so that data can be delivered faster.
Forword Proxy:

    A forward proxy sits between the client and the internet/server.
    client sends the request to the proxy, and then the proxy sends it to the destination.
    The destination server sees the proxy's IP, rather than directly seeing the client's IP in the network connection.  

    Example COmpany network uses proxy:
     - Control which websites employees can access
     - Block certain websites
     - filter traffic according to company policy
     - Apply security rules  
     - Cache some content      

Reverse Proxy:

    A reverse proxy is a server that sits in front of one or more backend servers and receives requests from users on behalf of servers.

    - Hide backend servers
    - Route requests
    - Caching
    - Authentication
    - Load balancing

Caching:

    caching means temporarily storing frequently used data so that we can retrieve it faster instead of fetching it from the original source(DB) every time.
    Redis: Redis is a software program that ues RAM to store/retrive data, that's why it faster.

 API Gateway:

     The API Gateway routes each request to the appropriate service.

     Responsibility: Routing, Authentication, Rate limiting ....
                  
             
        

