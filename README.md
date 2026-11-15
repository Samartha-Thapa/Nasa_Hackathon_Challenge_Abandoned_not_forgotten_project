# Welcome to Defold

This project was created from the "empty" project template.

The settings in ["game.project"](defold://open?path=/game.project) are all the default. A bootstrap empty ["main.collection"](defold://open?path=/main/main.collection) is included.

Check out [the documentation pages](https://defold.com/learn) for examples, tutorials, manuals and API docs.

If you run into trouble, help is available in [our forum](https://forum.defold.com).

Happy Defolding!

## Opening sequence

On startup, the game plays a short space flyby, reveals the animated Earth,
and introduces Commander Hayes before handing control to the player. Press
`E` or click after the briefing to enter the Earth base. Movement is disabled
until the briefing is dismissed.

The opening is isolated in `main/earth_opening_sequence.collection`: stars,
planets, asteroids, the Terran Earth rotation, the high-resolution surface
approach, and the commander briefing are layered independently before the
playable `main/main.collection` is loaded. The Earth stays at the same
150-pixel scale as the surrounding planets during the space shot, while the
high-resolution Earth is used only for the smooth zoom and base reveal.

---