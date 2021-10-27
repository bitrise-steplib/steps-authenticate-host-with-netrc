# Authenticate host with netrc

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/steps-authenticate-host-with-netrc?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/steps-authenticate-host-with-netrc/releases)

Adds your authentication configuration to the `.netrc` file.

<details>
<summary>Description</summary>

[This Step](https://github.com/bitrise-steplib/steps-authenticate-host-with-netrc) adds the authentication configuration (host name, login name and password string) to the `.netrc` file .
The Step lets you store your remote credentials on the build VM once so that later steps can use the credentials for authentication instead of requiring manual input. Examples include HTTPS git clone URLs with OAuth token-based authentication (instead of authenticating with SSH key).
Please note that if you already have a `.netrc` file, the Step will create a backup of the original, and appends the configs to the current one.

### Configuring the Step
1.Add the **Host** name, where the username and password will be used, for example, github.com.
2.Add the **Username**.
3.Add the password or the authentication token/ access token in the respective field which will be used by the host to authenticate you.

### Useful links
- [Learn more what the .netrc file format comprises of](https://everything.curl.dev/usingcurl/netrc#the-netrc-file-format)

### Related Steps
- [Activate SSH key (RSA private key)](https://www.bitrise.io/integrations/steps/activate-ssh-key)
- [Connect to OpenVPN Server](https://www.bitrise.io/integrations/steps/flutter-installer)
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `host` | The host where the username and password will be used. For example: github.com | required |  |
| `username` | The username used for the host to authenticate. | required, sensitive |  |
| `password` | The password (or Auth Token/Access Token) used for the host to authenticate. | required, sensitive |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-authenticate-host-with-netrc/pulls) and [issues](https://github.com/bitrise-steplib/steps-authenticate-host-with-netrc/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
