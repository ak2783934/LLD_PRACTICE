# 🎯 Requirements (clarify with interviewer first)

- Users (buyers) place orders with multiple items.
- Orders can have multiple statuses (Created → Confirmed → Shipped → Delivered → Cancelled).
- Inventory should be updated atomically.
- If multiple users order at the same time, stock consistency must be maintained.
- Notifications should go to buyers/suppliers on status updates.
- Extensible → easy to add new order states, new inventory types, etc.


