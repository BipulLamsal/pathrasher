## Ray Tracer

Rays are fired from the camera, through each pixel, into the scene
• When a ray hits an object, you compute a single surface point
• From that point, the renderer sends shadow rays toward each light source
• Those shadow rays are just visibility tests:
– if the light is visible → add its contribution
– if something blocks it → the light contributes nothing

• No light energy is bouncing around the scene
• No surface-to-surface illumination
• No indirect lighting
• No color bleeding

• Reflection and refraction rays may exist, but:
– they are limited in depth
– they still only gather direct lighting at each hit

So when the book says:

“the ray is then split into multiple rays that each go to a different light source”

What it means is:
• the renderer spawns multiple shadow rays
• one per light
• purely to check illumination
• not because light is physically splitting


## Diffuse Stuff (Matte Surfaces)
When light hits something matte (like a ball or paper), it doesn't just bounce off perfectly like a mirror. It scatters in random directions.
To simulate this, we:
1. Find where the ray hits.
2. Pick a random direction (somewhere on the hemisphere around the normal).
3. Shoot a new ray that way.
4. Darken the color a bit (like 50%) to show that some light got absorbed.

## Gamma Correction (Why is it dark?)
Screens are weird. If you tell a monitor to show "50% brightness" (0.5), it actually shows something that looks way darker (like 20% grey) to our eyes.
If we don't fix this, our whole render looks muddy and crushed.
**The Fix**: We "gamma correct" it. Basically, we take the square root of the color before saving it. This brightens up the mid-tones so they actually look right on screen.

## Shadow Acne (The speckled dots bug)
Computers aren't perfect at math. Sometimes when we calculate a hit point, the float value is just *slightly* off—like, 0.0000001 units inside the sphere.
If we shoot the next ray from there, it immediately hits the *same* sphere again and thinks it's blocked. This causes ugly black dots everywhere (shadow acne).
**The Hack**: We just ignore any hits that are super close to zero (anything less than 0.001). This lets the ray "escape" the surface.

## Lambertian vs Uniform (Making it look real)
* **Uniform**: Just picking a totally random direction. It works, but it's not quite how real physics works.
* **Lambertian**: This is the upgraded version. We pick a direction that favors pointing "straight up" from the surface (using strict probability). This mimics how real objects scatter light more accurately.



