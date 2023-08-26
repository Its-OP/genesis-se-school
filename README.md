# Rate Manager
Rate Manager is a tool that allows users to stay updated with the latest exchange rates of cryptocurrencies, specifically Bitcoin (BTC) to Ukrainian Hryvnia (UAH). It utilizes APIs of Binance, Coinbase, and Bitfinex to fetch real-time exchange rate data and provides a subscription feature for users to receive email notifications when the exchange rate updates.

Key Features:
- Real-time Exchange Rate: The project utilizes APIs from multiple providers to fetch the current BTC to UAH exchange rate. It ensures that users receive up-to-date information on the cryptocurrency market.
- Subscription Management: Users can subscribe to the service by providing their email addresses.
- Email Notifications: Subscribers receive email notifications whenever there is a change in the BTC to UAH exchange rate. The project leverages SendGrid to send personalized emails containing the updated exchange rate information.

# Get Started
1. Clone the repository;
2. Run `docker compose up` at the root folder of the project;
3. Whenever the compose stack launches, go to http://localhost:8080/swagger/index.html to access the service responsible for processing the currency data, or http://localhost:8080/swagger/index.html to access the service responsible for the email campaigns.

# !!! Important
**The application relies on SendGrid for sending emails. In order to enable sending emails by the application, you need to provide details of your SendGrid sender to the application, using environmental variables.**

**The next information needs to be passed to the application:**
1. Sender's name, via `SENDGRID_API_SENDER_NAME`.
2. Sender's email via `SENDGRID_API_SENDER_EMAIL`.
3. API key via `SENDGRID_API_KEY`.

**Information about your senders can be taken from this link:** https://app.sendgrid.com/settings/sender_auth/senders.
**The API key can be taken from here:** https://app.sendgrid.com/settings/api_keys.

**Placeholders for all these variables have been added to `docker-compose`. If you use it to run the application, please set the real values instead of the placeholders.**
