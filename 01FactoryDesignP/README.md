The Factory Design Pattern is a creational design pattern used to create objects without exposing the creation logic to the client.

👉 Instead of using new (or struct initialization) directly,\
👉 You delegate(saupna) object creation to a factory method.

Benefits (Backend Perspective)

✅ Loose coupling\
✅ Follows Open/Closed Principle (SOLID)\
✅ Easy to add new types\
✅ Cleaner controller/service logic\

Where You’ll Use It in Real Backend

-Creating DB connections (MySQL / PostgreSQL).\
-Creating payment gateways.\
-Creating notification services (email / SMS / push).
