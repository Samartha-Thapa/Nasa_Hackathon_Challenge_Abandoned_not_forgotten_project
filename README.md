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

The briefing uses `assets/sound_effects/typing_sound_effect.mp3` converted to
the Defold-compatible `typing_sound_effect.ogg`. It is attached through
`main/opening/typing.sound`, starts when the typewriter text appears, and stops
when the message finishes or the briefing is skipped.

After the briefing, Commander Hayes gives a guided first objective: walk to the
highlighted Lot 1 by following the animated dashed route. At the lot, a
blinking on-screen `OPEN SHOP` button opens the construction store.
The store remains hidden during normal gameplay and opens only beside a lot.
It presents one focused item card at a time with Previous/Next controls and
left/right swipe navigation, plus Buildings, Resources, Upgrades, and Close
controls. The first building card is the guided BUILD action; resources and
upgrades are available as separate store sections after that. The build
controller still validates costs, limits, lot state, and spawned building
factories; the store only presents the interaction and sends purchase requests.
The route guide follows the central access-road approach with a turn at the
outpost. Its cyan dash markers are world-space sprites in
`main/route_guidance.collection`, so they stay anchored to the map instead of
sliding with the HUD as the camera follows the player. The keyboard interaction
shortcut is disabled; the visible shop button is the intended entry point.
The focused card uses nine-slice framing so its border remains clean at
different sizes. Building cards use the existing buildings atlas and resource
cards use the `Space_ships_assets` frames registered in the store atlas.

---