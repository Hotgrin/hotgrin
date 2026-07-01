# An action that can fail, handled with try / if it fails.
action safe divide with a, b
    if b is 0
        give back problem "cannot divide by zero"
    end if
    give back a divided by b
end action

try
    set good to safe divide with 10, 2
    say "10 / 2 = " plus good
    set bad to safe divide with 10, 0
    say "you will never see this: " plus bad
if it fails
    say "Caught a problem: " plus the problem
end try

say "the program keeps running afterwards"
