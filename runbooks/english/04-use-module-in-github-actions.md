# Use module in GitHub Actions

Now we have the components of our CI pipeline, we're going to integrate them into a GitHub Action to automate our pipeline.

To do this, you can use the official [Dagger GitHub Action](https://docs.dagger.io/integrations/github).

> [!TIP]
> Dagger integrates with many existing tools and CI systems.
>
> You can see the [list on Dagger documentation](https://docs.dagger.io/integrations)

Create a new git branch (please replace `<branch_name>`):

```bash
git checkout -b <branch_name>
```

> [!NOTE]
> Add your GitHub username in the branch name to avoid duplicate branch.

Open `.github/workflows/CI.yaml` file in VSCode (you can find it on left panel).

Update `CI hello` GitHub Action - using the official Dagger GitHub Actions - to call our `Publish` function to publish the application.

To test the GitHub Action, push your branch and create a Pull Request.

You now have a CI pipeline for the `hello` application!

The pipeline is executed on GitHub Action using our Dagger module! This allows you to use the same pipeline locally and in GitHub Actions.

To go further, you can [Create, publish a module in Daggerverse and use it](05-create-publish-module.md).
