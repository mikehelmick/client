# Are you here for the fun(k)?

This branch contains the demo for the proposed [fun(k)](https://docs.google.com/document/d/1V7M32ELv38SvHciYBloYlibRSs_e8i2ACqnYL6h5YEo/edit#) (Functions on Knative).

Here are the instructions for running the demo. Sorry, it's not quite plug and
play yet, we wanted to get something out for people to see quickly. If this gets
picked up as part of Knative, the first actions will be to make this "just
work."

The demo only work with Go at the moment. Should this work move forward, nodejs
will be prioritized. And, as the SDK spec will be open... whatever else you
can imagine.

## fun(k) Setup

First - this assumes you have a Kubernetes cluster with Knative Serving AND
Knative Eventing installed and setup with a `Broker` available in the `default`
namespace.

1. Clone this fork to `$GOPATH/src/knative.dev/client`
2. Change to that directory and checkout the funk branch, `git checkout funk`
3. Build, `./hack/build.sh`
4. Optional - create an alias for `kn` to the version you just built. Just make
sure that if you have the binary distro of kn installed, make sure that you
running this version when you go through the commands below.
5. Create a funk settings directory at `~/.kn/funk`
6. Clone the go SDK repo from https://github.com/mikehelmick/funk-sdk-go to
`$GOPATH/src/knative.dev/funk-sdk-go`
7. Edit the `sdks.json` file in the `funkData` directory so that the absolute
path for the funk-sdk-go (line 9) directory matches where you cloned that repo
to in step 5.
8. Copy the sdks.json to the right location.
`cp funkData/sdks.json ~/.kn/funk/sdks.json`
9. Apply the provided EventType `kubectl apply -f funkData/event_type.yaml` -
This uses a protype public schema registry hosted at
https://schemas.in-the-cloud.dev
  * If you want to do custom types, the schema you reference in the `EventType`
    just needs to be a valid URI
10. The funk-go-sdk is currently hardcoded to work with a single client project
location. Make a directory at `$GOPATH/src/github.com/mikehelmick/play`

## fun(k) demo

Still here? Cool, here's how you run the demo.
For this, in the current state, you also need to have the following installed

* https://github.com/golang/dep
* https://github.com/google/ko
  * ko needs to be configured to point to your docker repository correctly

1. `cd $GOPATH/src/github.com/mikehelmick/play`
2. `kn service list`
  * Just to make sure you're connected to your cluster with `kn`
3. `kn funk`
  * To show that you have the right build, with kn funk present
4. `kn funk languages list`
  * you should see the Go SDK and the (fake) NodeJS one
5. `kn funk init go`
6. `kn funk function create sayHello --t com.mikehelmick.eventutils.user.user`
  * When this is done, the function is in `pkg/funk`
7. `kn funk deploy`
8. [optional] Edit the function - and deploy again with `kn funk deploy`

To generate an HTTP function

1. `kn funk function create helloWorld --http=true`
2. `kn funk deploy`

## Other things that are used in the demo

1. https://github.com/mikehelmick/ceproxy - Simple CloudEvents proxy that
takes events from off cluster and puts them on the specified broker.
2. https://github.com/mikehelmick/eventutils - The first cut of code generation
  * the cmd/typegen program can be used to generate schemas from Go types

# Knative Client

This section outlines best practices for the Knative developer experience. It is a reference for Knative CLI implementation, and a reference for Knative client libraries.

The goals of the Knative Client are to:

1. Follow the Knative [serving](https://github.com/knative/serving) and [eventing](https://github.com/knative/eventing) APIs
2. Be scriptable to allow users to create different Knative workflows
3. Expose useful Golang packages to allow integration into other programs or CLIs or plugins
4. Use consistent verbs, nouns, and flags for various commands
5. Be easily extended via a plugin mechanism (similar to `kubectl`) to allow for experimentation and customization

# Docs

Start with the [user's guide](docs/README.md) to learn more. You can read about common use cases, get detailed documentation on each command, and learn how to extend the `kn` CLI. For more information, access the following links:

* [User's guide](docs/README.md)
* [Generated documentation](docs/cmd/kn.md)

**Shell auto completion:**

Run the following command to enable shell auto-completion:

For Zsh:
```sh
$ source <(kn completion zsh)
```

For Bash:
```sh
$ source <(kn completion bash)
```

Use TAB to list available sub-commands or flags.

# Developers

If you would like to contribute, please see
[CONTRIBUTING](https://knative.dev/contributing/)
for more information.

To build `kn`, see our [Development](DEVELOPMENT.md) guide.
