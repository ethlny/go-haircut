export type ApiResponse = {
    code: number;
    status: string;
    message: string;
    data: any;
};

export const login = async (email: string, password: string, type: string): Promise<ApiResponse> => {
    const res = await fetch(`/api/auth/${type}_login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password })
    });

    const data: ApiResponse = await res.json();
    return data;
}

type PossibleBody = {
    email: string;
    password: string;
    name?: string;
    address?: string;
    phone?: string;
    openingTime?: string;
    closingTime?: string;
}

export const signup = async (body: PossibleBody, type: string): Promise<ApiResponse> => {
    if (type === 'client') {
        const res = await fetch(`/api/auth/${type}_signup`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ email: body.email, password: body.password })
        });

        const data: ApiResponse = await res.json();
        return data;
    }
    else if (type === 'saloon') {
        const res = await fetch(`/api/auth/${type}_signup`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body)
        });

        const data: ApiResponse = await res.json();
        return data;
    } else {
        return {
            code: 400,
            status: 'error',
            message: 'Invalid type',
            data: null
        }
    }
}
