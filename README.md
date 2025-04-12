# Database Backup Utility
A Database Backup Utility is a command-line tool designed in **Go** to simplify the process of backing up and restoring databases. It supports various database systems, including MySQL, PostgreSQL.

Github URL: [https://github.com/sambasivareddy-ch/db_backup_utility](https://github.com/sambasivareddy-ch/db_backup_utility)

## Table of Contents
1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Folder Structure](#folder-structure)
5. [Feature Roadmap](#feature-roadmap)
6. [Contributing](#contributing)
7. [License](#license)
8. [Acknowledgements](#acknowledgements)
9. [Contact Information](#contact-information)

## Features
+ Supports multiple database systems:
    - MySQL
    - PostgreSQL
+ Backup and restore functionality
+ Command-line interface for easy usage
+ Error handling and logging
+ Discord Notification for backup & restore status
+ Configurable settings for different environments
+ Allow Schedule backup using Cron jobs.

## Installation 
Requirements:
- Go 1.22.3 or later
- Git
- Database client (MySQL/PostgreSQL) installed
- Discord webhook URL for notifications

### Step 1: Clone the repository
Git clone the repository to your local machine:
```bash
    git clone  https://github.com/sambasivareddy-ch/db_backup_utility.git
```
### Step 2: Change directory
```bash
    cd db_backup_utility
```
### Step 3: Install dependencies
```bash
    go mod tidy
```
### Step 4: Set up environment variables
Create a `.env` file in the root directory of the project and add your Discord webhook URL:
```bash
    DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/your_webhook_url
```
### Step 5: Build the project
```bash
    go build -o db_backup_utility main.go
```
or 
```bash
    go build -o db_backup_utility .
```

## Usage 
### Get help
To get help on how to use the utility, run the following command:
```bash
    ./db_backup_utility help
```
### Configure the utility
To configure the utility with necessary database connection details and backup directory details 
```bash
    ./db_backup_utility config
```
It will prompt you to enter the following details:
- Database type (MySQL/PostgreSQL)
- Database host
- Database port
- Database name
- Database user
- Database password
- Backup directory
### Backup a database
To backup a database, run the following command:
```bash
    ./db_backup_utility backup
```
This will create a backup of the specified database and save it in the configured backup directory.
In case of MySQL, it will create a `mysql_backup_<backup_time>.sql` file with the backup.
In case of PostgreSQL, it will create a `pg_backup_<backup_time>.dump` file with the backup.
### Restore a database
To restore a database from a backup, run the following command:
```bash
    ./db_backup_utility restore
```
This will prompt you with list of available backups in the configured backup directory.
You can select a backup file to restore the database.
### Schedule backups
To schedule backups using cron jobs, you can use the following command:
```bash
    ./db_backup_utility schedule -c <cron_expression> -o <operation>
```
Replace `<cron_expression>` with the desired cron expression (e.g., `0 2 * * *` for daily backups at 2 AM) and `<operation>` with either `backup` or `restore`.

## Folder Structure
```bash
    ├── cmd
    │   ├── backup.go
    │   ├── config.go
    │   ├── restore.go
    │   ├── schedule.go
    │   └── root.go
    ├── context
    │   └── context.go
    ├── pkg
    │   ├── backup
    │   │   ├── postgres.go
    │   │   └── mysql.go
    │   ├── cron
    │   │   ├── cron.go
    │   ├── db
    │   │   ├── mysql.go
    │   │   └── postgres.go
    │   ├── executor
    │   │   ├── executor.go
    │   ├── utils
    │   │   ├── dir_check.go
    │   │   └── interactive_terminals.go
    │   ├── logging
    │   │   ├── logger.go
    │   ├── notify
    │   │   ├── notify.go
    ├── main.go
    ├── .env
    ├── .gitignore
    ├── go.mod
    ├── go.sum
    ├── LICENSE
    ├── app.log
    ├── db_context.json
    └── README.md
```

## Feature Roadmap
- [ ] Add support for additional database systems (e.g., SQLite, MongoDB)
- [ ] Implement encryption for backup files
- [ ] Add more advanced backup options (e.g., incremental backups, compression)
- [ ] Add support for cloud storage (e.g., AWS S3, Google Cloud Storage)

## Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements
This project was inspired by the need for efficient database management tools and the contributions of the open-source community.

## Contact Information
### For any questions or feedback, please contact
Email Address: sambachinta.24@gmail.com <br>
LinkedIn: [Sambasiva Reddy Chinta](https://www.linkedin.com/in/samba-siva-reddy-ch/)
