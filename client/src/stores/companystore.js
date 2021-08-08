import { writable } from "svelte/store"

export const mcompany = writable([])
const fetchCompany = async() => {
    const url = `http://localhost:4000/api/company`
    const res = await fetch(url)
    const data = await res.json()
    const loadedCompany = data.record.map((data, index) => {
        return {
            no : index + 1,
            idcompany: data.idcompany,
            idcurr: data.idcurr,
            nmcompany: data.nmcompany,
            nmowner: data.nmowner,
            statuscompany: data.statuscompany,
        }
    })
    mcompany.set(loadedCompany)
}
fetchCompany()