# testerr

# Introduction
The testerr package provides a simple, data driven method of testing errors in
data driven tests. When a table driven test needs to include tests which
inspect an error that may be returned, instead of embedding an error and flags
to control how to test it or just hardcoding the error handling into the test,
the testerr package allows the data table for the tests to include the expected
error and how to check it.

# License
The testerr package is licensed under the MIT license. Please see the LICENSE
file for details.