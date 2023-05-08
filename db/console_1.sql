alter table delivery add constraint fk_delivery_id
foreign key (delivery_id) references order_table(delivery_id);

alter table payment add constraint fk_payment_id
foreign key (payment_id) references order_table(payment_id);

alter table items add constraint fk_order_id
foreign key (order_id) references order_table(orderuid);