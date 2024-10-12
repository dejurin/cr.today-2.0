# How to language update by goi18n?

## Run commands:

1. goi18n extract -sourceLanguage en -outdir localize .
2. goi18n merge -outdir localize localize/active.en.toml localize/translate.uk.toml
3. goi18n merge -outdir localize active.uk.toml translate.uk.toml

