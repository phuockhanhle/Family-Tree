import { FormHelperText, makeStyles } from "@material-ui/core";

export const useStyles = makeStyles(() => ({

    toto: {
        display: 'flex',
        justifyContent: 'center',
    },

    pageView: {
        display: 'flex',
        justifyContent: 'center',
    },

    appView: {
        maxWidth: 1440,
    },

    navigationBar: {
        backgroundColor: 'white',
        position: 'relative',
        maxHeight: 80,
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        paddingLeft: 30,
        paddingRight: 30,
        '@media (max-width: 780px)' : {
            flexDirection: 'column',
            maxHeight: 'inherit',
        }
    },

    title: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },

    logo: {
        // paddingBottom: 0,
        paddingTop: 3,

    },

    line: {
        marginRight: 13,
        marginLeft: 3,
        paddingTop: 5,
    },

    titleName: {
        color: '#004687',
        fontSize: 30,
        cursor: 'pointer',
        '@media (max-width: 780px)' : {
           marginBottom: 10,
        }
    },

    account: {
        color: 'black',
        display: 'flex',
        alignItems: 'center',
    },

    accountName: {
        color: '#004687',
        fontWeight: 'bold',
        marginRight: 10,
        fontSize: 18,
        textTransform: 'none',
        cursor: 'pointer',
    },

    popOver: {
        "& .MuiPaper-rounded": {
            borderRadius: 20,
        }
    },

    popOverItem: {
        '&:hover': {
            backgroundColor: '#E8E8E8',
        },
    },

    questionMark: {
        marginRight: 10,
        cursor: 'pointer',
    },

    logOut: {
        cursor: 'pointer',
    },

    container: {
        paddingTop: 10,
        maxWidth: '100vw',
        maxHeight: '100%',
        backgroundColor: '#F4F4F4',
        '@media (max-width: 780px)' : {
            flexDirection: 'row',
         }
    },

    communicationZone: {
        backgroundColor: '#fff',
        borderRadius: 20,
        marginTop: 25,
        marginBottom: 25,
    },

    communicationInnerZone: {
        display: 'flex',
    },

    communicationWarning: {
        color: 'white',
        width: 65,
        textAlign: 'center',
        backgroundColor: 'red',
        borderTopLeftRadius: 20,
        borderBottomLeftRadius: 20,
        "+ p": {
            paddingLeft: 20,
        }
    },

    communicationMessage: {
        paddingLeft: 20,
    },

    ContentCol: {
        flexDirection: 'row',
        '@media (max-width: 780px)' : {
            flexDirection: 'column-reverse',
         }
    },

    ColG: {
        maxWidth: '75%',
        flexBasis: '75%',
        '@media (max-width: 780px)' : {
            maxWidth: '100%',
            flexBasis: '100%',
         }
    },

    ColD: {
        maxWidth: '25%',
        flexBasis: '25%',
        '@media (max-width: 780px)' : {
            maxWidth: '100%',
            flexBasis: '100%',
         }
    },

    details: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        maxHeight: 78,
    },

    ActusLink: {
        color: '#0078BE',
        cursor: 'pointer',
        fontSize: 18,
        marginTop: 0,
        textAlign: 'right',
        fontFamily: 'FFF Equipe',
        fontWeight: 'bold',
    },

    Outils: {
        backgroundColor: 'white',
        borderRadius: 20,
        padding: 20,
        marginTop: 25,
    },

    who: {
        color: '#014687',
        fontFamily: 'FFF Equipe',
        fontWeight: 'bold',
        fontSize: 30,
        "& > span:first-child": {
            color: '#3D3D3D',
            fontWeight: '100',
        },
        '@media (max-width: 780px)' : {
            fontSize: 20,
            margin: 0,
         }
    },

    where: {
        color: '#014687',
        fontSize: 18,
        '@media (max-width: 780px)' : {
            fontSize: 16,
            marginBottom: 15,
         }
    },

    divider: {
        backgroundColor: '#014687',
        marginBottom: 50,
    },

    vignetteZone: {
        backgroundColor: 'white',
        borderRadius: 20,
        padding: 20,
        paddingBottom: 40,
        marginBottom: 30,
        marginRight: 30,
    },

    Actualite: {
        backgroundColor: 'white',
        borderRadius: 20,
        padding: 20,
    },

    vignetteZoneTitle: {
        color: '#014687',
        fontSize: 30,
        // fontWeight: 'bold',
        textTransform: 'uppercase',
        marginTop: 0,
        marginBottom: 0,
        fontFamily: 'FFF Hero',
    },

    underLine: {
        marginTop: 0,
        marginBottom: 30,
        display: 'flex',
        alignItems: 'top',
    },

    ContentVign: {
        padding: '10px !important',
    },

    vignetteBox: {
        backgroundImage: 'url("../static/LogoPictosEtFonds/FondVignettes-BlocFOOT2000.png")',
        filter: 'drop-shadow(0px 0px 10px rgba(0, 0, 0, 0.1))',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        display: 'block',
        borderRadius: 20,
        border: 'solid 1px transparent',
        minHeight: 120,
        maxWidth: 297,
        // transition: 'box-shadow 0.15s ease-out, transform 0.25s ease',
        '&:hover': {
            // transform: 'scale(1.1)',
            cursor: 'pointer',
        },
        '@media (max-width: 780px)' : {
            display: 'inline-flex',
            minHeight: 105,
            width: '100%',
        }
    },

    vignetteBoxBlue: {
        backgroundImage: 'url("../static/LogoPictosEtFonds/FondVignettes-BlocServicesAdministratifs.png")',
        filter: 'drop-shadow(0px 0px 10px rgba(0, 0, 0, 0.1))',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        display: 'block',
        borderRadius: 20,
        border: 'solid 1px transparent',
        minHeight: 120,
        // transition: 'box-shadow 0.15s ease-out, transform 0.25s ease',
        '&:hover': {
            // transform: 'scale(1.1)',
            cursor: 'pointer',
        },
        '@media (max-width: 780px)' : {
            display: 'inline-flex',
            minHeight: 105,
            width: '100%',
        }
    },

    vignetteBoxWhite: {
        backgroundImage: 'url("../static/LogoPictosEtFonds/FondVignettes-BlocRéseauBleuetUniversFFF.png")',
        filter: 'drop-shadow(0px 0px 10px rgba(0, 0, 0, 0.1))',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        display: 'block',
        border: 'solid 1px transparent',
        borderRadius: 20,
        minHeight: 120,
        // transition: 'box-shadow 0.15s ease-out, transform 0.25s ease',
        '&:hover': {
            // transform: 'scale(1.1)',
            cursor: 'pointer',
        },
        '@media (max-width: 780px)' : {
            display: 'inline-flex',
            minHeight: 105,
            width: '100%',
        }
    },

    vignetteIcon: {
        display: 'flex',
        justifyContent: 'center',
        marginTop: 30,
        flexDirection: 'column',
        alignItems: 'center',
        '@media (max-width: 780px)' : {
            margin: 0,
            padding: 0,
            width: '100%',
            alignItems: 'center',
        }
    },

    vignetteTextWhite: {
        display: 'flex',
        justifyContent: 'center',
        textAlign: 'center',
        marginTop: 6,
        color: '#FFFFFF',
        fontWeight: 'bold',
        fontSize: 20,
        '@media (max-width: 500px)' : {
            fontSize: 11,
        },
        '@media (min-width: 500px)' : {
            fontSize: 14,
        },
        '@media (min-width: 780px)' : {
            fontSize: 20,
        },
    },

    vignetteTextBlue: {
        display: 'flex',
        justifyContent: 'center',
        textAlign: 'center',
        marginTop: 6,
        color: '#014586',
        fontWeight: 'bold',
        fontSize: 20,
        '@media (max-width: 780px)' : {
            fontSize: 14,
        }
    },
}));