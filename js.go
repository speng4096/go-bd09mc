package bd09mc

const script = `
var b6 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";

function bN(cE) {
    var cC = "";
    var cL, cJ, cH = "";
    var cK, cI, cG, cF = "";
    var cD = 0;
    var T = /[^A-Za-z0-9\+\/\=]/g;
    if (!cE || T.exec(cE)) {
        return cE
    }
    cE = cE.replace(/[^A-Za-z0-9\+\/\=]/g, "");
    do {
        cK = b6.indexOf(cE.charAt(cD++));
        cI = b6.indexOf(cE.charAt(cD++));
        cG = b6.indexOf(cE.charAt(cD++));
        cF = b6.indexOf(cE.charAt(cD++));
        cL = (cK << 2) | (cI >> 4);
        cJ = ((cI & 15) << 4) | (cG >> 2);
        cH = ((cG & 3) << 6) | cF;
        cC = cC + String.fromCharCode(cL);
        if (cG != 64) {
            cC = cC + String.fromCharCode(cJ)
        }
        if (cF != 64) {
            cC = cC + String.fromCharCode(cH)
        }
        cL = cJ = cH = "";
        cK = cI = cG = cF = ""
    } while (cD < cE.length);
    return cC
}

function bV(T) {
    return typeof T == "string"
}

function b4(T, cC) {
    if (isNaN(T)) {
        T = bN(T);
        T = isNaN(T) ? 0 : T
    }
    if (bV(T)) {
        T = parseFloat(T)
    }
    if (isNaN(cC)) {
        cC = bN(cC);
        cC = isNaN(cC) ? 0 : cC
    }
    if (bV(cC)) {
        cC = parseFloat(cC)
    }
    this.lng = T;
    this.lat = cC
}

var map = {
    EARTHRADIUS: 6370996.81,
    MCBAND: [12890594.86, 8362377.87, 5591021, 3481989.83, 1678043.12, 0],
    LLBAND: [75, 60, 45, 30, 15, 0],
    MC2LL: [[1.410526172116255e-8, 0.00000898305509648872, -1.9939833816331, 200.9824383106796, -187.2403703815547, 91.6087516669843, -23.38765649603339, 2.57121317296198, -0.03801003308653, 17337981.2], [-7.435856389565537e-9, 0.000008983055097726239, -0.78625201886289, 96.32687599759846, -1.85204757529826, -59.36935905485877, 47.40033549296737, -16.50741931063887, 2.28786674699375, 10260144.86], [-3.030883460898826e-8, 0.00000898305509983578, 0.30071316287616, 59.74293618442277, 7.357984074871, -25.38371002664745, 13.45380521110908, -3.29883767235584, 0.32710905363475, 6856817.37], [-1.981981304930552e-8, 0.000008983055099779535, 0.03278182852591, 40.31678527705744, 0.65659298677277, -4.44255534477492, 0.85341911805263, 0.12923347998204, -0.04625736007561, 4482777.06], [3.09191371068437e-9, 0.000008983055096812155, 0.00006995724062, 23.10934304144901, -0.00023663490511, -0.6321817810242, -0.00663494467273, 0.03430082397953, -0.00466043876332, 2555164.4], [2.890871144776878e-9, 0.000008983055095805407, -3.068298e-8, 7.47137025468032, -0.00000353937994, -0.02145144861037, -0.00001234426596, 0.00010322952773, -0.00000323890364, 826088.5]],
    LL2MC: [[-0.0015702102444, 111320.7020616939, 1704480524535203, -10338987376042340, 26112667856603880, -35149669176653700, 26595700718403920, -10725012454188240, 1800819912950474, 82.5], [0.0008277824516172526, 111320.7020463578, 647795574.6671607, -4082003173.641316, 10774905663.51142, -15171875531.51559, 12053065338.62167, -5124939663.577472, 913311935.9512032, 67.5], [0.00337398766765, 111320.7020202162, 4481351.045890365, -23393751.19931662, 79682215.47186455, -115964993.2797253, 97236711.15602145, -43661946.33752821, 8477230.501135234, 52.5], [0.00220636496208, 111320.7020209128, 51751.86112841131, 3796837.749470245, 992013.7397791013, -1221952.21711287, 1340652.697009075, -620943.6990984312, 144416.9293806241, 37.5], [-0.0003441963504368392, 111320.7020576856, 278.2353980772752, 2485758.690035394, 6070.750963243378, 54821.18345352118, 9540.606633304236, -2710.55326746645, 1405.483844121726, 22.5], [-0.0003218135878613132, 111320.7020701615, 0.00369383431289, 823725.6402795718, 0.46104986909093, 2351.343141331292, 1.58060784298199, 8.77738589078284, 0.37238884252424, 7.45]],
    getDistanceByMC: function (cG, cE) {
        if (!cG || !cE) {
            return 0
        }
        var cC, cF, T, cD;
        cG = this.convertMC2LL(cG);
        if (!cG) {
            return 0
        }
        cC = this.toRadians(cG.lng);
        cF = this.toRadians(cG.lat);
        cE = this.convertMC2LL(cE);
        if (!cE) {
            return 0
        }
        T = this.toRadians(cE.lng);
        cD = this.toRadians(cE.lat);
        return this.getDistance(cC, T, cF, cD)
    },
    getDistanceByLL: function (cG, cE) {
        if (!cG || !cE) {
            return 0
        }
        cG.lng = this.getLoop(cG.lng, -180, 180);
        cG.lat = this.getRange(cG.lat, -74, 74);
        cE.lng = this.getLoop(cE.lng, -180, 180);
        cE.lat = this.getRange(cE.lat, -74, 74);
        var cC, T, cF, cD;
        cC = this.toRadians(cG.lng);
        cF = this.toRadians(cG.lat);
        T = this.toRadians(cE.lng);
        cD = this.toRadians(cE.lat);
        return this.getDistance(cC, T, cF, cD)
    },
    convertMC2LL: function (cC) {
        var cD, cF;
        cD = new b4(Math.abs(cC.lng), Math.abs(cC.lat));
        for (var cE = 0; cE < this.MCBAND.length; cE++) {
            if (cD.lat >= this.MCBAND[cE]) {
                cF = this.MC2LL[cE];
                break
            }
        }
        var T = this.convertor(cC, cF);
        var cC = new b4(T.lng.toFixed(6), T.lat.toFixed(6));
        return cC.lng + '|' + cC.lat
    },
    convertLL2MC: function (T) {
        var cC, cE;
        T.lng = this.getLoop(T.lng, -180, 180);
        T.lat = this.getRange(T.lat, -74, 74);
        cC = new b4(T.lng, T.lat);
        for (var cD = 0; cD < this.LLBAND.length; cD++) {
            if (cC.lat >= this.LLBAND[cD]) {
                cE = this.LL2MC[cD];
                break
            }
        }
        if (!cE) {
            for (var cD = this.LLBAND.length - 1; cD >= 0; cD--) {
                if (cC.lat <= -this.LLBAND[cD]) {
                    cE = this.LL2MC[cD];
                    break
                }
            }
        }
        var cF = this.convertor(T, cE);
        var T = new b4(cF.lng.toFixed(2), cF.lat.toFixed(2));
        return T.lng + '|' + T.lat
    },
    convertor: function (cD, cE) {
        if (!cD || !cE) {
            return
        }
        var T = cE[0] + cE[1] * Math.abs(cD.lng);
        var cC = Math.abs(cD.lat) / cE[9];
        var cF = cE[2] + cE[3] * cC + cE[4] * cC * cC + cE[5] * cC * cC * cC + cE[6] * cC * cC * cC * cC + cE[7] * cC * cC * cC * cC * cC + cE[8] * cC * cC * cC * cC * cC * cC;
        T *= (cD.lng < 0 ? -1 : 1);
        cF *= (cD.lat < 0 ? -1 : 1);
        return new b4(T, cF)
    },
    getDistance: function (cC, T, cE, cD) {
        return this.EARTHRADIUS * Math.acos((Math.sin(cE) * Math.sin(cD) + Math.cos(cE) * Math.cos(cD) * Math.cos(T - cC)))
    },
    toRadians: function (T) {
        return Math.PI * T / 180
    },
    toDegrees: function (T) {
        return (180 * T) / Math.PI
    },
    getRange: function (cD, cC, T) {
        if (cC != null) {
            cD = Math.max(cD, cC)
        }
        if (T != null) {
            cD = Math.min(cD, T)
        }
        return cD
    },
    getLoop: function (cD, cC, T) {
        while (cD > T) {
            cD -= T - cC
        }
        while (cD < cC) {
            cD += T - cC
        }
        return cD
    }
};`
