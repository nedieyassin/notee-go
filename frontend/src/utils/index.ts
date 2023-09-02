export const formatDateString = (dateString: number): string => {
    const date = new Date(dateString)
    const now = new Date()

    const timeDifference = now.getTime() - date.getTime()
    const withinOneDay = timeDifference < 24 * 60 * 60 * 1000
    const withinOneYear = now.getFullYear() - date.getFullYear() < 1

    const hours = date.getHours().toString().padStart(2, "0")
    const minutes = date.getMinutes().toString().padStart(2, "0")
    const day = date.getDate().toString().padStart(2, "0")
    const month = (date.getMonth() + 1).toString().padStart(2, "0")
    const year = date.getFullYear()

    if (withinOneDay) {
        return `${hours}:${minutes}`
    } else if (withinOneYear) {
        return `${hours}:${minutes} ${day}/${month}/${year}`
    } else {
        return `${day}/${month}/${year}`
    }
}

export const formatNumber = (num?: number): string => {
    if (!num) {
        return '--'
    }
    if (num >= 1000000) {
        return (num / 1000000).toFixed(1) + 'M';
    } else if (num >= 1000) {
        return (num / 1000).toFixed(1) + 'K';
    } else {
        return num.toString();
    }
}

export const appendNumberToTitle = (title: string, titles: string[]): string => {
    let newTitle = title;
    let count = 1;

    while (titles.includes(newTitle)) {
        newTitle = `${title} ${count}`;
        count++;
    }

    return newTitle;
}