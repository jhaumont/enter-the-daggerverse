# Use module in github actions

Now we have the components of our CI pipeline, we're going to integrate them into a GitHub action to automate our pipeline.

To do this, you can use the official [Dagger's GitHub Action](https://docs.dagger.io/integrations/github) integration.

> [!TIP]
> Many Dagger integrations exist for many tools and CI. 
>
> You can see the [list on Dagger documentation](https://docs.dagger.io/integrations)

Create a new git branch (please replace `<nom_branche>`):

```bash
git checkout -b <nom_branche>
```

> [!NOTE]
> Add your GitHub nickname in the branch name to avoid duplicate branch.

Open `.github/workflows/CI.yaml` file in VSCode (you can find it on left panel).

Update `CI hello` GitHub Action - using official Dagger's GitHub Actions - to call our `Publish` function to publish the application.

To test the GitHub Action, push your branch and create a Pull Request.

You now have a CI pipeline for `hello` application.

The pipeline is executed on GitHub Action with using our Dagger's functions!

To go further, you can [Create, publish a module in Daggerverse and use it](05-create-publish-module.md).
