import { fetchUtils, withLifecycleCallbacks, DataProvider } from 'react-admin';
import simpleRestProvider from './myDataProvider';

type CloudinaryFile = {
    asset_id: string;
    key: string;
};
type SignData = {
    token: string;
    domain: string;
}

function generateSecureRandomString(length: number): string {
    const array = new Uint8Array(length);
    window.crypto.getRandomValues(array);
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    return Array.from(array, byte => characters[byte % characters.length]).join('');
}


const fetchJson = (url: string, options: any = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    // add your own headers here
    const token = localStorage.getItem('token');
    if (token) {
        options.headers.set('Authorization', `Bearer ${token}`);
    }
    return fetchUtils.fetchJson(url, options);
}
export interface PackageItem {
    package: string;
    func: FuncItem[];
    id: number;
}

export interface FuncItem {
    name: string;
    alias: string;
    description: string;
    active: boolean;
}

const dataProvider = withLifecycleCallbacks(
    simpleRestProvider(import.meta.env.VITE_SIMPLE_REST_URL, fetchJson),
    [
        {
            resource: "article",
            beforeSave: async (params: any, dataProvider: DataProvider) => {
                if (params.picture.rawFile) {
                    const response = await fetch(
                        `${import.meta.env.VITE_SIMPLE_REST_URL}/auto/upload`,
                        {
                            method: "GET", headers: {
                                "Content-Type": "application/json",
                                "Authorization": `Bearer ${localStorage.getItem('token')}`
                            }
                        }
                        // should send headers with correct authentications
                    );

                    const signData: SignData = await response.json();
                    console.log(signData);
                    const url = `https://up-z2.qiniup.com`;

                    const formData = new FormData();
                    formData.append("file", params.picture.rawFile);
                    formData.append("token", signData.token);
                    let fileName = params.picture.rawFile.name || generateSecureRandomString(10)
                    formData.append("key", fileName);
                    // formData.append("bucket", "ginblog");
                    const imageResponse = await fetch(url, {
                        method: "POST",
                        body: formData,
                    });

                    const image: CloudinaryFile = await imageResponse.json();
                    return {
                        ...params,
                        picture: {
                            src: `${signData.domain}${image.key}`,
                            title: image.key,
                        },
                    };
                }

                return params;
          
            },
        },
        {
            resource: "systemMenu",
            afterGetList: async (data, dataProvider) => {
                let idCounter = 1; // ID 计数器
                (data as any).data.forEach((packageItem: PackageItem) => {
                    packageItem['id'] = idCounter++
                    packageItem.func.forEach((func: any) => {
                        func['active'] = false
                    })
                })
                return data
            }
        }
    ]
)


export default dataProvider;