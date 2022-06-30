<script>
	import { onMount } from 'svelte';
	import * as d3 from 'd3';
	import { palettes } from "./colors.js"
	import { settings } from './store.js';

	export let data;
	console.log(data);

	// TODO create weekday graph
	console.log(data);

	let viz;
	let margin = {top: 10, right: 30, bottom: 30, left: 60},
    width = 580 - margin.left - margin.right,
    height = 300 - margin.top - margin.bottom;

	onMount(() => {
		let svg = d3.select(viz)
			.append("svg")
    		.attr("width", width + margin.left + margin.right)
  			.attr("height", height + margin.top + margin.bottom)
			.append("g")
			.attr("transform", "translate(" + margin.left + "," + margin.top + ")");

		let xScale = d3.scaleBand()
			.domain(data.map(d => d.weekday))
			.range([0, width])
			.padding(0.25);
			// .tickValues(data) // add for custom values
			// .tickFormat(d => (d.year +":"+ d.val));

		svg.append("g")
			.attr("transform", "translate(0," + height + ")")
			.call(d3.axisBottom(xScale));
			// .selectAll("text")  
            // .style("text-anchor", "end")
            // .attr("dx", "-.8em")
            // .attr("dy", ".15em")
            // .attr("transform", "rotate(-65)" );

		let yScale = d3.scaleLinear()
			.domain([0, d3.max(data, d => d.minutes)])
			.range([ height, 0 ]);

    	svg.append("g")
    		.call(d3.axisLeft(yScale));

		// let line = d3.line()
		// 	.x(d => xScale(d.date))
		// 	.y(d => yScale(d.minutes))

		// svg.append("path")
		// 	.datum(data)
		// 	.attr("fill", "none")
		// 	.attr("stroke", palettes[$settings.theme].primary)
		// 	.attr("stroke-width", 2.5)
		// 	.attr("d", line);

		svg.selectAll("bar")
			.data(data)
			.enter()
			.append("rect")
			.attr("x", d => xScale(d.weekday))
			.attr("y", d => yScale(d.minutes))
			.attr("width", xScale.bandwidth())
			.attr("height", d => height - yScale(d.minutes))
			.attr("fill", palettes[$settings.theme].primary);
	});
</script>

<div bind:this={viz} class="chart"></div>

<style>
	div {
		margin: 3em 0;
	}
</style>