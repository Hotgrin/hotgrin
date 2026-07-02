# A calculator with tests written in plain English.
action add with a, b
    give back a plus b
end action

action discount with price, percent
    give back price minus (price times percent divided by 100)
end action

test "addition works"
    expect add with 2, 3 to be 5
    expect add with 10, 10 to be 20
end test

test "ten percent off 100 is 90"
    expect discount with 100, 10 to be 90
end test

test "discount never goes above the price"
    expect discount with 50, 10 to be at most 50
end test
