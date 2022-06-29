// theme color sets
export const palettes = {
    "Default": {
        bg: "#fff",
        primary: "#1b4965",
        hover: "#bee9e8",
        press: "#62b6cb",
        fgCircle: "#62b6cb",
        bgCircle: "#bee9e8",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },
    "Purp": {
        bg: "#e2dcf1",
        primary: "#4c3388",
        hover: "#baacdb",
        press: "#4c3388",
        fgCircle: "#4c3388",
        bgCircle: "#baacdb",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },
    "Cafe": {
        bg: "#F4DFBA",
        primary: "#876445",
        hover: "#cfb27e",
        press: "#876445",
        fgCircle: "#876445",
        bgCircle: "#CA965C",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },
    "Cream": {
        bg: "#ECE5C7",
        primary: "#ada390",
        hover: "#C2DED1",
        press: "#ada390",
        fgCircle: "#C2DED1",
        bgCircle: "#e3dcbc",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#444",
        textOnPrimary: "#fff",
    },
    "Slate": {
        bg: "#eeeeee",
        primary: "#aaaaaa",
        hover: "#bbbbbb",
        press: "#cccccc",
        fgCircle: "#bbbbbb",
        bgCircle: "#dddddd",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#333",
        textOnPrimary: "#fff",
    },
    "Eve": {
        bg: "#1B2430",
        primary: "#51557E",
        hover: "#816797",
        press: "#51557E",
        fgCircle: "#816797",
        bgCircle: "#2f3e52",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#fff",
        textOnPrimary: "#fff",
    },
    "Cold": {
        bg: "#b7cce0",
        primary: "#53717f",
        hover: "#7394a1",
        press: "#53717f",
        fgCircle: "#5f95ad",
        bgCircle: "#7394a1",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },

    "Pastel": {
        bg: "#fdffb6",
        primary: "#bdb2ff",
        hover: "#ffc6ff",
        press: "#a0c4ff",
        fgCircle: "#ffadad",
        bgCircle: "#ffd6a5",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },
    "Little": {
        bg: "#F9D923",
        primary: "#187498",
        hover: "#EB5353",
        press: "#36AE7C",
        fgCircle: "#187498",
        bgCircle: "#36AE7C",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },

    "Pumkin": {
        bg: "#1B1A17",
        primary: "#E45826",
        hover: "#6e554d",
        press: "#E45826",
        fgCircle: "#E45826",
        bgCircle: "#6e554d",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#fff",
        textOnPrimary: "#fff",
    },
    "Pink": {
        bg: "#fbc4ab",
        primary: "#f08080",
        hover: "#f8ad9d",
        press: "#f4978e",
        fgCircle: "#f4978e",
        bgCircle: "#f8ad9d",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#333",
        textOnPrimary: "#fff",
    },
    "Lettuce": {
        bg: "#FCF9C6",
        primary: "#297312",
        hover: "#b3c482",
        press: "#6da82d",
        fgCircle: "#6da82d",
        bgCircle: "#b3c482",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#000",
        textOnPrimary: "#fff",
    },
    "Contrast": {
        bg: "#000",
        primary: "#fff",
        hover: "#aaa",
        press: "#fff",
        fgCircle: "#fff",
        bgCircle: "#333",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#fff",
        textOnPrimary: "#000",
    },
    "Night": {
        bg: "#051923",
        primary: "#003554",
        hover: "#006494",
        press: "#0582ca",
        fgCircle: "#0582ca",
        bgCircle: "#006494",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#aaa",
        textOnPrimary: "#aaa",
    },
    "Ocean": {
        bg: "#a9d6e5",
        primary: "#2c7da0",
        hover: "#61a5c2",
        press: "#89c2d9",
        fgCircle: "#1da5e0",
        bgCircle: "#89c2d9",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#333",
        textOnPrimary: "#fff",
    },
    "Shore": {
        bg: "#fee9e1",
        primary: "#64b6ac",
        hover: "#c0fdfb",
        press: "#64b6ac",
        fgCircle: "#64b6ac",
        bgCircle: "#fad4c0",
        error: "#f00",
        errorHover: "#fbb", 
        textOnBG: "#333",
        textOnPrimary: "#fff",
    },
};

export let setCSS = (theme) => {
    const palette = palettes[theme];

    const cssStr = `
        --bg: ${palette.bg};
        --primary: ${palette.primary};
        --hover: ${palette.hover};
        --press: ${palette.press};
        --fg-circle: ${palette.press};
        --bg-circle: ${palette.bgCircle};
        --error: ${palette.error};
        --error-hover: ${palette.errorHover}; 
        --text-on-bg: ${palette.textOnBG};
        --text-on-primary: ${palette.textOnPrimary};
    `;

    document.documentElement.style.cssText = cssStr;
}
