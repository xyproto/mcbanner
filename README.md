# mcbanner

Linux application for generating Minecraft banners.

## Includes

* A web server/page showing the variations of Minecraft banners, as SVG approximations (`mcweb`).
* A commandline application for evolving Minecraft banners with GAs (`evolve`).
* Go code for generating a random SVG banner and rendering it as PNG by using `rsvg-convert` (`random`).

## The original goal

* Given an image, get the steps for creating the closest looking Minecraft banner. This goal is a work in progress, since the evolution is not rapid enough. Perhaps using smaller images when rendering from SVG, or another algorithm, would work.

## General information

* Version: 0.1
* License: MIT
* Author: Alexander F. Rødseth
