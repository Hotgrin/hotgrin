# A tiny shop calculation.
describe order
    item is "Wireless mouse"
    price is 299
    quantity is 3
end describe

set total to price of order times quantity of order
say "Order: " plus item of order
say "Total: R" plus total

action discount with amount, percent
    give back amount minus (amount times percent divided by 100)
end action

say "After 10% off: R" plus discount with total, 10
