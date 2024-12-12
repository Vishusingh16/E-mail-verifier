const apiURL ='http://localhost:8080/api'; // We have to replace this with out golang backend server

export const signup = async(email, password )=>{
    try{
        const response = await fetch(`${apiURL}/signup`,{
            method : 'POST',
        headers :{'Content-Type':'application/json'},
        body: JSON.stringify({email, password}),

        });
        return await response.json();

    }catch (err){
        console.error('API error :' ,err);

    }

};

export const login = async(email, password)=>{
    try{
        const response = await fetch('${apiURL}/login',{
            method: 'POST',
            headers :{'Content-Type': 'application/json'},
            body: JSON.stringify({email, password}),
        });
        return await response.json();

    }catch(err){
        console.error('API error :', err);
    }
};

