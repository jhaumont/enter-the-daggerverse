# Use module in github actions

Now that we have all components for our CI, we will include it in Github Action to automate our pipeline.

Todo this, you can use official [Dagger's GitHub Action](https://docs.dagger.io/integrations/github)

> [!TIP]
> Many Dagger integration exists in many tools and CI. 
>
> You can see the lastest [list on Dagger documentation](https://docs.dagger.io/integrations)

Create a new git branch (please replace `<nom_branche>`):


```bash
git checkout -b <nom_branche>
```

> [!NOTE]
> Add you nickname GitHub in branch name to avoid double branch.

Open `.github/workflows/CI.yaml` file in VSCode (you can find it un left panel).

Update `CI hello` GitHub Action - with using official Dagger's GitHub Actions - to call our `Publish` function to publish our application.

To test the GitHub Action, push your branch and create a Pull Request.

You have now a CI pipeline for `hello` application.

The pipeline is execution on GitHub Action with using our Dagger's functions!

To the next, please go to the next page [Create, publish a module in Daggerverse and use it](05-create-publish-module.md).
