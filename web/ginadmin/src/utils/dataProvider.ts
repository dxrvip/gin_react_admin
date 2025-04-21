import { fetchUtils, withLifecycleCallbacks, DataProvider } from 'react-admin';
// import simpleRestProvider from './myDataProvider';
import simpleRestProvider from 'ra-data-simple-rest';

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
    return fetchUtils.fetchJson(url, options).then(response => {
        const { json } = response
        // console.log(json.data)
        // 如果没有data
        if (json.code !== 200 || !json?.data) {
            throw new Error(json.message || "错误的请求体！")
        }
        return {
            status: response.status,
            headers: response.headers,
            body: response.body,
            json: json.data,
        };

    });
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
// 新增上传相关工具函数
const getUploadToken = async () => {
    const response = await fetch(
        `${import.meta.env.VITE_SIMPLE_REST_URL}/auto/upload`,
        {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem('token')}`
            }
        }
    );
    return await response.json() as SignData;
};

interface RawFileData {
    rawFile: File;
}

interface ImageData {
    src: string | RawFileData;
    title: string;
}


interface UploadParams {
    picture?: ImageData[];
    [key: string]: any;
}

const createQiniuUpload = async (file: File, signData: SignData): Promise<CloudinaryFile> => {
    try {
        const formData = new FormData();
        formData.append("file", file);
        formData.append("token", signData.token);
        const fileName = file?.name || generateSecureRandomString(10);
        formData.append("key", fileName);

        const response = await fetch("https://up-z2.qiniup.com", {
            method: "POST",
            body: formData,
        });

        if (!response.ok) {
            throw new Error(`Upload failed: ${response.statusText}`);
        }

        return await response.json();
    } catch (error) {
        console.error('Upload error:', error);
        throw error;
    }
};

const uploadImage = async (params: UploadParams, dataProvider: DataProvider): Promise<UploadParams> => {
    if (!Array.isArray(params.picture)) {
        return params;
    }

    try {
        const uploaded = await Promise.all(
            params.picture.map(async (pic: ImageData) => {
                // 处理新上传的图片
                if (pic && typeof pic.src === 'string' && 'rawFile' in pic) {
                    const signData = await getUploadToken();
                    const image = await createQiniuUpload(pic.rawFile as File, signData);
                    return {

                        src: `${signData.domain}${image.key}`,
                        title: pic.title || image.key,

                    };
                }

                // 如果是已有的图片数据，转换为正确格式
                if (pic?.src && typeof pic.src === 'string') {
                    return {

                        src: pic.src,
                        title: pic.title || ''

                    };
                }

                // 如果已经是正确的格式，直接返回
                if (pic?.src && typeof pic.src === 'string') {
                    return pic;
                }

                return null; // 无效的图片数据
            })
        );
        const nonNullUploaded = uploaded.filter(item => item !== null) as ImageData[];


        return {
            ...params,
            picture: nonNullUploaded,
        };
    } catch (error) {
        console.error('Image processing error:', error);
        throw error;
    }
};

const dataProvider = withLifecycleCallbacks(
    simpleRestProvider(import.meta.env.VITE_SIMPLE_REST_URL, fetchJson),
    [
        {
            resource: "article",
            beforeSave: async (params: UploadParams, dataProvider: DataProvider) => {
                return await uploadImage(params, dataProvider);
            },
        },
        {
            resource: "product",
            beforeSave: async (params: UploadParams, dataProvider: DataProvider) => {
                return await uploadImage(params, dataProvider);
            }
        },
        {
            resource: "secondHandSkus",
            beforeSave: async (params: UploadParams, dataProvider: DataProvider) => {
                return await uploadImage(params, dataProvider);
            }
        },
        {
            resource: "systemMenu",
            afterGetList: async (data, dataProvider) => {
                let idCounter = 1;
                (data as any).data.forEach((packageItem: PackageItem) => {
                    packageItem['id'] = idCounter++;
                    packageItem.func.forEach((func: any) => {
                        func['active'] = false;
                    });
                });
                return data;
            }
        }
    ]
);


export default dataProvider;