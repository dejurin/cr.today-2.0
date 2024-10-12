## Translating Messages with Go-i18n

If you have added new messages to your program and want to handle translation updates, follow the steps below to ensure all active and translated message files are up to date.

Step-by-Step Guide

1. Extract new messages

When you add new messages to your program, you need to extract these messages to the active locale file. This will ensure that the base language (usually English) contains all the messages.

Run the following command to extract new messages and update active.en.toml:

goi18n extract -sourceLanguage en -outdir localize .

This command scans your code for new messages and adds them to active.en.toml.

2. Generate updated translation files

Next, you will generate the translation files (translate.*.toml) for all languages by merging the new messages into these files.

Run the following command:

goi18n merge -outdir localize localize/active.*.toml

This command updates the translation files, such as translate.uk.toml, by adding any new messages that are not yet translated. The new messages will be marked for translation.

3. Translate new messages

Open the relevant translation file (e.g., translate.uk.toml) and manually translate the newly added messages. Example of a translated entry:

["Currency.UAH"]
few = "{{.count}} гривні"
many = "{{.count}} гривень"
one = "гривня"
other = "{{.count}} гривні"

Translate all the new messages in this file.

4. Merge translated messages

Once you have finished translating the new messages, merge the translations back into the active locale files. This will ensure that your application can use the updated translations.

Run the following command:

goi18n merge -outdir localize localize/active.*.toml localize/translate.*.toml

This will merge the translations from translate.uk.toml (or other translation files) into active.uk.toml (or other active locale files), creating a complete set of translated messages for your application.

5. Using the updated message files

Now that the messages have been merged, the updated active.*.toml files can be used in your application to provide localized content.

Notes:

	•	Make sure you keep both active.*.toml and translate.*.toml files under version control to track changes.
	•	If you add new languages, you will need to repeat this process for the new language’s translate file.
