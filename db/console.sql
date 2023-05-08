create table Delivery
(
    Delivery_id text primary key,
    Name        text,
    Phone       text,
    Zip         text,
    City        text,
    Adress      text,
    Region      text,
    Email       text
);

create table Payment
(
    Payment_id   text primary key,
    Transaction  text,
    RequestId    text,
    Currency     text,
    Provider     text,
    Amount       int,
    PaymentDt    int,
    Bank         text,
    DeliveryCost int,
    GoodsTotal   int,
    CustomFee    int
);

create table Items
(
    Order_id    text,
    Items_id    text primary key,
    ChrtId      int,
    TrackNumber text,
    Price       int,
    Rid         text,
    Name        text,
    Sale        int,
    Size        text,
    TotalPrice  int,
    MnId        int,
    Brand       text,
    Status      int
);

create table Order_table
(
    OrderUid          text primary key,
    TrackNumber       text,
    Entry             text,
    Delivery_id       text,
    Payment_id        text,
    Locale            text,
    InternalSignature text,
    CustomerId        text,
    DeliveryService   text,
    Shardkey          text,
    SmId              int,
    DateCreated       timestamp,
    OofShard          text
);