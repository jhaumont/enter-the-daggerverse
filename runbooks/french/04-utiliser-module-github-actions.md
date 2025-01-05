# Utiliser le module dans une GitHub Actions

Maintenant que nous avons les composants de notre pipeline de CI, nous allons les intégrer à une GitHub action pour automatiser notre pipeline.

Pour se faire, il existe une integration officielle de Dagger pour GitHub actions : https://docs.dagger.io/integrations/github

> [!TIP]
> Il existe un grand nombre d'integration Dagger dans divers outils de build & CI.
>
> La liste exhaustive se trouve ici: https://docs.dagger.io/integrations

Créer une nouvelle branche Git (remplacer `<nom_branche>`) :

```bash
git checkout -b <nom_branche>
```

> [!NOTE]
> Ajoutez votre pseudo GitHub dans le nom de votre branche pour éviter les doublons.

Ouvrir le fichier `.github/workflows/CI.yaml` dans VSCode (à trouver dans le panel Explorer sur la gauche).

Modifier la GitHub action `CI hello` - en utilisant l'intégration Dagger pour GitHub Actions - afin d'appeler la fonction `Publish` pour builder et publier l'application.

Pour tester la GitHub Actions, pousser votre branche et créer une Pull Request.

Vous avez maintenant un pipeline de CI pour l'application `hello`.

Le pipeline s'exécute sur GitHub Actions en utilisant les fonctions Dagger que vous avez créées.

Pour aller plus loin, vous pouvez [créer, publier un module dans le daggerverse et l'utiliser](05-créer-oublier-utiliser-module.md) dans un pipeline.
