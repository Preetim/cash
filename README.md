## Cash coding exercise

### Running tests instructions

There are 2 ways you can run this code;
1. Docker:
You need to install docker and run this command
 `docker build .`
   
2. Using your local setup:
[Install go](https://golang.org/doc/install),
   once done run `go test` in your project folder

   
### Design Concepts

This bank is made of accounts and transactions. Accounts have names and amounts for each customer. Every time a deposit or withdraw is done, the record is written in as a transaction and then the amount is reflected in the customer's account.
Transactions form a kind of ledger for the bank. 
In an ideal world the bank balance calculation and transactions would be events which would run as scheduled to create settlement files.

### Data Model 
Data model is very simple as shown below.

![img_1.png](img_1.png)

### Trade Offs/Assumptions
- All amounts are represented in cents and hence the type is integers.
- No concept of time or currency or statuses shown for this code.
- The id fields in the data model arent in this code. They'd be primary keys in a DB.
- I havent handled concurrency in this code. I'd be implementing optimistic locking in the DB with manual retries
- Would implement auto-retries for transient failures like app not available, maybe with the help of queues and helpers
- Transactions would be backed up regularly in case the primary DB fails to ensure no money is lost for the bank


Feedback

#### Feedback for candidate

Thank you for your submission. Your solution was simple and concise which made it easy to understand. Your ReadMe was good, although we would have loved to see your reasoning behind the trade offs and assumptions. We felt that there were some edge cases missing, such as checking for negative amounts in deposits and withdrawals. Technical Design could have used some improvement, consider how the code will be used outside the package by specifying a public API.

Overall, good job!

[+] **ReadMe**:

- Instructions on how to run the tests were clear and easy to follow, however some further details on how to install dependencies would be helpful
- Good list of assumptions, however some discussion on the reasoning behind the assumptions would have been good

[+] **Testing** :

- Test suite is generally clear and complete
- ~~There were a couple of asserts that could have been more specific (for example TestWithdrawalWithInsufficientBalance checks the balance is not -10, when asserting the balance was still 50 would have been testing the actual amount, or asserting the full error message, rather than there simply being an error)~~
- ~~The TestPrintSettlementFile test could be more useful if it had some assertions but it only prints data.
  findBankBalance hasn’t been tested thoroughly with different deposit/withdrawal operations.~~

[+] **Analysis**:

- Withdraw and deposit operations do not check that the account argument exists (i.e. not null).
- ~~Interestingly, the findAccount method returns a default “empty” account instance. The implication being that withdrawing/depositing from an account that doesn’t exist would end up with the bank applying the operation on the default account which is error prone~~.
- There’s no validation against negative amounts.
- Otherwise meets all requirements

[-] **Technical Design**:

- ~~The bank implementation is not reusable from a different package to the one where the bank sits. All methods in the bank start with a lowercase letter which makes them package-private according to Go standards.~~
- ~~Tests are modifying internal state directly rather than doing this through the exposed interface~~
- ~~Tests modify the accounts slice directly instead of using an exposed operation.~~
- ~~The accounts slice is not reflected automatically when an account is created. calculateBankBalance relies on this slice being populated.~~
- ~~The transaction slice is populated regardless of whether the account actually exists. This abstraction also seems to be unnecessary as it was not mentioned in the requirements and is not used throughout the solution.~~

[+] **Code Fluency**:

- The candidate could’ve used table-driven testing to increase coverage.
- The solution uses idiomatic Go.

[-] **Code Maintainability**:

- Generally clear and concise code.
- Correctness of the solution could’ve been better.
- Submitted JetBrains files.
- ~~Redundant comments on the test scenarios.~~

#### Did the candidate reach an acceptable solution?
Yes
