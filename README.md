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
  * the cmd/generate program can be used to generate Clients and Functions
    (this is the basis for fun(k))

# Kn

The Knative client `kn` is your door to the [Knative](https://knative.dev)
world. It allows you to create Knative resources interactively from the command
line or from within Shell scripts.

`kn` offers you:

- Full support for managing all features of
  [Knative Serving](https://github.com/knative/serving) (services, revisions,
  traffic splits)
- Growing support [Knative eventing](https://github.com/knative/eventing),
  closely following its development (managing of sources & triggers)
- A plugin architecture similar to that of `kubectl` plugins
- A thin client-specific API in golang which helps in tasks like synchronously
  waiting on Knative service write operations.
- An easy integration of Knative into Tekton Pipelines by using
  [`kn` in a Tekton `Task`](https://github.com/tektoncd/catalog/tree/master/kn).

This client uses the
[Knative Serving](https://github.com/knative/docs/blob/master/docs/serving/spec/knative-api-specification-1.0.md)
and
[Knative Eventing](https://github.com/knative/eventing/tree/master/docs/spec)
API exclusively so that it will work with any Knative installation, even those
that are not Kubernetes based. It does not help in _installing_ Knative itself
though. Please refer to the various
[Knative installation options](https://knative.dev/docs/install/) for how to
install Knative with its prerequisites.

## Documentation

Start with the [user's guide](docs/README.md) to learn more. You can read about
common use cases, get detailed documentation on each command, and learn how to
extend the `kn` CLI. For more information, have a look at:

- [User guide](docs/README.md)
  - Installation - How to install `kn` and run on your machine
  - Examples - Use case based examples
  - FAQ (_to come._)
- [Reference Manual](docs/cmd/kn.md) - all possible commands and options with
  usage examples

## Developers

We love contributions! Please refer to
[CONTRIBUTING](https://knative.dev/contributing/) for more information on how to
best contributed to contribute to Knative.

For code contributions it as easy as picking an
[issue](https://github.com/knative/client/issues) (look out for
"kind/good-first-issue"), briefly comment that you would like to work on it,
code, test, code and finally submit a
[PR](https://github.com/knative/client/pulls) which will trigger the review
process.

More details on how to build and test can be found in the
[Developer guide](docs/DEVELOPMENT.md).
