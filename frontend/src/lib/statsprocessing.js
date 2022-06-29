import { isoParse } from "d3"; 

// TODO make algorithm more effecient
// TODO move to server side processing

export function getGraphData(pomodoros) {
    const now = new Date();

	const pomos = pomodoros.map(d => {
		return {
			date: isoParse(d.timestamp),
			minutes: d.minutes
		}
	});

    return {
        week: daysPast(pomos, 7, now),
        month: daysPast(pomos, 30, now),
        threemonths: daysPast(pomos, 90, now),
    }
}

const day = 24 * 60 * 60 * 1000;

function daysPast(pomos, n, now) {
    let data = new Array(n);
    const weekPast = now.getTime() - n * day;
    let filtered = pomos.filter(d => d.date.getTime() > weekPast);

    for (let i = n-1; i >= 0; i--) {
        let time = now.getTime() - i*day;
        let currDate = new Date(time).toDateString();
        let currMinutes = filtered.reduce((prev, next) => {
            if (next.date.toDateString() == currDate) {
                return prev + next.minutes;
            } else {
                return prev;
            }
        }, 0);

        data[n-1-i] = {
            date: currDate,
            minutes: currMinutes,
        };
    }

    console.log(data);

    return data;
}